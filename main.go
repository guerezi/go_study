package main

import (
	"context"
	"imobiliaria/internal/repositories/mysql"
	"imobiliaria/internal/usecases"
	"imobiliaria/internal/validators"
	"imobiliaria/server"
	"imobiliaria/server/handlers"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/sethvargo/go-envconfig"
	"github.com/sirupsen/logrus"
)

// Deus está triste com tanto copilot

/// air

const DefaultPort = ":3000"

type Config struct {
	Host     string `env:"DB_HOST" default:"localhost"`
	Port     string `env:"DB_PORT" default:"3306"`
	User     string `env:"DB_USER" default:"root"`
	Password string `env:"DB_PASSWORD" default:"password"`
	Database string `env:"DB_NAME" default:"database"`
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	logrus.SetLevel(logrus.TraceLevel)

	ctx := context.Background()

	// Configurações do banco de dados
	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		panic(err)
	}

	// obg chat gpt
	// Preicsa ser ponteiro pq to recendo interface (??)
	m, err := mysql.NewRepository(&mysql.Config{
		Host:     c.Host,
		Port:     c.Port,
		User:     c.User,
		Password: c.Password,
		Database: c.Database,
	})

	if err != nil {
		logrus.WithError(err).Fatal("Error creating repository")
	}

	// TODO: validator cuistomizado iaaaa
	v := validator.New()

	rules := map[string]string{
		"street":  `^[A-Za-zÀ-ÿ0-9\sºª.,'-]{3,100}$`,
		"number":  `^[0-9]{1,6}[A-Za-z\-ºª\/]{0,10}$`,
		"city":    `^[A-Za-zÀ-ÿ\s'-]{2,100}$`,
		"state":   `^[A-Z]{2}$`,
		"zipcode": `^[0-9]{5}-[0-9]{3}$`,
	}

	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		logrus.WithError(err).Fatal("Error registering default translations")
	}

	for index, val := range rules {
		if err := v.RegisterValidation(index, validators.RegexValidator(val)); err != nil {
			logrus.WithError(err).Fatalf("Error registering %s validator", index)
		}

		err = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
			return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())

			return t
		})

		if err != nil {
			logrus.WithError(err).Fatalf("Error registering %s translation", index)
		}
	}

	u := usecases.NewUsecases(m, v)
	h := handlers.Handler{
		Usecases:  u,
		Validator: v,
	}

	/// H com & quer dizer que não é uma copia
	s := &server.Server{
		Handler: &h,
	}

	if err := s.Listen(DefaultPort); err != nil {
		logrus.WithError(err).Fatal("Server error")
	}
}
