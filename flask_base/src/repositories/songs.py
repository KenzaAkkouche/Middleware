from src.helpers import db
from src.models.song import Song

def get_all_songs(self):
    return Song.query.all()

def get_song_from_id(id):
    return Song.query.get(id)

def add_song(song):
    db.session.add(song)
    db.session.commit()

def update_song(song):
    existing_song = get_song_from_id(song.id)
    existing_song.Title = song.Title
    existing_user.Artist = song.Artist
    db.session.commit()

def delete_song(id):
    song = get_song_from_id(id)
    if song:
        db.session.delete(song)
        db.session.commit()
        return True
    return False

