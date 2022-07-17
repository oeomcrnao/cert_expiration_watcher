# About

A project for tracking the expiration of ssl certificates and sending notifications
 (currently only Telegram is available)

## Setup

To use it, you need to set the following environment variables:

- TARGET_DOMAIN - Domain to check
- TARGET_PORT - Port to check
- TELEGRAM_ENABLED - Enable send notification in Telegram
- TELEGRAM_BOT_TOKEN - Telegram bot token
- TELEGRAM_CHAT_ID - Telegram chat id (you can get it via @userinfobot)
- TRESHOLD_IN_DAYS - How many days before the expiration of the domain you need to send a notification

## Usage

```docker build . -t your_image_name```

```docker run -e ... (environment variables) your_image_name```

OR

```docker run -e ... (environment variables) oeomcrnao/domain_expiration_watcher```

You can also use a file with environment variables

```docker run --env-file .env oeomcrnao/domain_expiration_watcher```

## Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement". Don't forget to give the project a star! Thanks again!

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.
