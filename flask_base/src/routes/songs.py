import json
from flask import Blueprint, request, jsonify
from marshmallow import ValidationError

from src.models.http_exceptions import *
from src.schemas.song import SongSchema
from src.schemas.errors import *
from src.schemas.song import SongUpdateSchema
import src.services.songs as song_service


songs = Blueprint(name="songs", import_name=__name__)



@songs.route('/', methods=['GET'])
def get_songs():
    songs_data, status_code = song_service.get_all_songs()
    return jsonify(SongSchema(many=True).dump(songs_data)), status_code


@songs.route('/<id>', methods=['GET'])
def get_song(id):
    return song_service.get_song(id)


@songs.route('/', methods=['POST'])
def create_song():
    try:
        song_ajout = SongSchema().load(request.get_json())
        return song_service.create_song(song_ajout)
    except ValidationError as e:
        error = UnprocessableEntitySchema().load({"message": str(e)})
        return error, error.get("code")
    except (Conflict, UnprocessableEntity, Forbidden) as e:
        error = e.get_error_schema()
        return error, error.get("code")
    except Exception:
        error = SomethingWentWrongSchema().load({})
        return error, error.get("code")

@songs.route('/<id>', methods=['PUT'])
def put_song(id):
    try:
        song_update = SongUpdateSchema().load(request.get_json())
        return song_service.put_song(id, song_update)
    except ValidationError as e:
        error = UnprocessableEntitySchema().load({"message": str(e)})
        return error, error.get("code")
    except (Conflict, UnprocessableEntity, Forbidden) as e:
        error = e.get_error_schema()
        return error, error.get("code")
    except Exception:
        error = SomethingWentWrongSchema().load({})
        return error, error.get("code")

@songs.route('/<id>', methods=['DELETE'])
def delete_song(id):
    try:
        return song_service.delete_song(id)
    except Exception as e:
        print(f"An error occurred: {e}")
        return jsonify({"message": "Internal Server Error"}), 500
