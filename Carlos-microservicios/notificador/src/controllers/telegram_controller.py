##!/usr/bin/env python
# -- coding: utf-8 --
#-------------------------------------------------------------------------
# Archivo: telegram_controller.py
# Capitulo: Estilo Microservicios
# Autor(es): Perla Velasco & Yonathan Mtz. & Jorge Solís
# Version: 3.0.0 Febrero 2022
# Descripción:
#
#   Ésta clase define el controlador del microservicio API. 
#   Implementa la funcionalidad y lógica de negocio del Microservicio.
#
#   A continuación se describen los métodos que se implementaron en ésta clase:
#
#                                             Métodos:
#           +------------------------+--------------------------+-----------------------+
#           |         Nombre         |        Parámetros        |        Función        |
#           +------------------------+--------------------------+-----------------------+
#           |     send_message()     |         Ninguno          |  - Procesa el mensaje |
#           |                        |                          |    recibido en la     |
#           |                        |                          |    petición y ejecuta |
#           |                        |                          |    el envío a         |
#           |                        |                          |    Telegram.          |
#           +------------------------+--------------------------+-----------------------+
#
#-------------------------------------------------------------------------
from flask import request, jsonify
import json
from src.helpers.config import load_config
import asyncio
from telegram import Bot

class TelegramController:

    @staticmethod
    def send_message():
        try:
            data = json.loads(request.data)
            message = data.get("message")

            if not message:
                return jsonify({"msg": "message field is required"}), 400

            config = load_config()
            token = config["telegram"]["token"]
            chat_id = config["telegram"]["chat_id"]
            bot = Bot(token=token)

            # Ejecutamos el envío de mensaje de forma asíncrona
            asyncio.run(bot.send_message(chat_id=chat_id, text=message))

            return jsonify({"msg": "Mensaje enviado a Telegram"}), 200

        except Exception as e:
            return jsonify({"msg": "Error al procesar la solicitud", "error": str(e)}), 500