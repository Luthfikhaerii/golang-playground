package messaging // package khusus untuk handle Kafka messaging

import (
	"context" // untuk kontrol lifecycle (cancel, timeout, shutdown)

	"github.com/IBM/sarama" // library Kafka client untuk Go
	"go.uber.org/zap"       // structured logging

	"golang-playground/pkg/logger" // custom logger project
)

// ConsumerHandler adalah function signature untuk business logic consumer
// menerima 1 message Kafka dan return error jika gagal
type ConsumerHandler func(message *sarama.ConsumerMessage) error

// ConsumerGroupHandler adalah adapter untuk memenuhi interface Sarama ConsumerGroupHandler
// membungkus function sederhana menjadi bentuk yang bisa dipakai Kafka
type ConsumerGroupHandler struct {
	Handler ConsumerHandler // function business logic yang akan dipanggil
}

// Setup dipanggil saat consumer mulai atau saat rebalance terjadi
// biasanya untuk init resource (sekarang belum dipakai)
func (h *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil // tidak ada aksi
}

// Cleanup dipanggil saat consumer berhenti atau sebelum rebalance ulang
// biasanya untuk cleanup resource (sekarang belum dipakai)
func (h *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil // tidak ada aksi
}

// ConsumeClaim adalah core function untuk membaca message dari Kafka
// function ini WAJIB ada karena bagian dari interface Sarama
func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for { // loop terus untuk consume message
		select {

		// case: menerima message dari Kafka (via channel)
		case message := <-claim.Messages():

			if message == nil { // jika tidak ada message (channel ditutup)
				return nil // keluar dari consumer
			}

			// jalankan business logic handler
			if err := h.Handler(message); err != nil {

				// jika gagal, log error
				// offset TIDAK di-mark → kemungkinan akan di-retry
				logger.Log.Error("failed to process message", zap.Error(err))

			} else {

				// jika sukses, tandai message sudah diproses
				// ini akan commit offset ke Kafka
				session.MarkMessage(message, "")
			}

		// case: jika context selesai (shutdown / rebalance)
		case <-session.Context().Done():
			return nil // keluar dari loop
		}
	}
}

// ConsumeTopic adalah helper untuk menjalankan Kafka consumer
// menghubungkan consumer group, topic, dan handler
func ConsumeTopic(ctx context.Context, consumerGroup sarama.ConsumerGroup, topic string, handler ConsumerHandler) {

	// bungkus handler function ke dalam adapter struct
	consumerHandler := &ConsumerGroupHandler{Handler: handler}

	// goroutine utama untuk consume Kafka
	go func() {
		for { // loop terus (penting untuk handle rebalance Kafka)

			// mulai consume topic
			if err := consumerGroup.Consume(ctx, []string{topic}, consumerHandler); err != nil {

				// log jika terjadi error saat consume
				logger.Log.Error("error from consumer", zap.Error(err))
			}

			// jika context sudah di-cancel (shutdown)
			if ctx.Err() != nil {
				logger.Log.Info("context cancelled, stopping consumer")
				return // keluar dari goroutine
			}
		}
	}()

	// goroutine untuk menangani error dari Kafka consumer group
	go func() {
		for err := range consumerGroup.Errors() {

			// log semua error dari Kafka
			logger.Log.Error("consumer group error", zap.Error(err))
		}
	}()

	// blocking sampai context di-cancel (misal: shutdown aplikasi)
	<-ctx.Done()

	// log bahwa consumer akan ditutup
	logger.Log.Info("closing consumer", zap.String("topic", topic))

	// tutup koneksi consumer group ke Kafka
	if err := consumerGroup.Close(); err != nil {

		// log jika gagal close
		logger.Log.Error("error closing consumer group", zap.Error(err))
	}
}
