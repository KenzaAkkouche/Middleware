import json
import requests
from marshmallow import EXCLUDE
from sqlalchemy import exc

from src.models.song import Song as SongModel
from src.schemas.song import SongSchema

import src.repositories.songs as songs_repository
from src.models.http_exceptions import SomethingWentWrong, Forbidden, UnprocessableEntity, Conflict

songs_url = "http://localhost:8081/songs/"


def get_all_songs():
    response = requests.request(method="GET", url=songs_url)
    return response.json(), response.status_code


def get_song(Id):
    response = requests.request(method="GET", url=songs_url+Id)
    return response.json(), response.status_code


def create_song(songs_register):
    # on récupère le modèle utilisateur pour la BDD
    song_model = SongModel.from_dict(songs_register)
    # on récupère le schéma utilisateur pour la requête vers l'API songs
    song_schema = SongSchema().loads(json.dumps(songs_register), unknown=EXCLUDE)

    # on crée l'utilisateur côté API songs
    response = requests.request(method="POST", url=songs_url, json=song_schema)
    if response.status_code != 200:
        return response.json(), response.status_code

    # on ajoute l'utilisateur dans la base de données
    # pour que les données entre API et BDD correspondent
    try:
        song_model.Id = response.json()["Id"]
        songs_repository.add_song(song_model)
    except Exception:
        raise SomethingWentWrong

    return response.json(), response.status_code

def put_song(Id, song_update):

    # s'il y a quelque chose à changer côté API
    song_schema = SongSchema().loads(json.dumps(song_update), unknown=EXCLUDE)
    response = None
    if not SongSchema.is_empty(song_schema):
        # on lance la requête de modification
        response = requests.request(method="PUT", url=songs_url+Id, json=song_schema)
        if response.status_code != 201:
            return response.json(), response.status_code

    # s'il y a quelque chose à changer côté BDD
    song_model = SongModel.from_dict(song_update)
    if not song_model.is_empty():
        song_model.Id = Id
        found_song = songs_repository.get_song_from_id(Id)
        if not song_model.Title:
            song_model.Title = found_song.Title
        if not song_model.Artist:
            song_model.Artist = found_song.Artist

        try:
            songs_repository.update_song(song_model)
        except exc.IntegrityError as e:
            if "NOT NULL" in e.orig.args[0]:
                raise UnprocessableEntity
            raise Conflict

    return (response.json(), response.status_code) if response else get_song(Id)

def delete_song(Id):
    response = requests.request(method="DELETE", url=songs_url + Id)

    if response.status_code == 204:
        return {"message": "Song deleted successfully"}, 200
    elif response.status_code == 404:
        return {"message": "Song not found"}, 404
    else:
        try:
            data = response.json()
            return data, response.status_code
        except json.JSONDecodeError:
            return {"message": "Invalid JSON received from Golang server"}, 500




def get_song_from_db(self, song_id):
    return self.songs_repository.get_song(song_id)

def song_exists(self, song_id):
    return self.get_song_from_db(song_id) is not None

