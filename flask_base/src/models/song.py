from src.helpers import db

class Song(db.Model):
    __tablename__ = 'songs'

    Id = db.Column(db.String(255), primary_key=True)
    Title = db.Column(db.String(255), nullable=False)
    Artist = db.Column(db.String(255), nullable=False)

    def __init__(self, uuid, Title, Artist):
        self.Id = uuid
        self.Title = Title
        self.Artist = Artist


    def is_empty(self):
        return (not self.Id or self.Id == "") and \
               (not self.Title or self.Title == "") and \
               (not self.Artist or self.Artist == "")

    @staticmethod
    def from_dict(obj):
        return Song(
            uuid=obj.get("uuid"),
            Title=obj.get("Title"),
            Artist=obj.get("Artist")
        )

# Ajout du modèle pour les évaluations (ratings)
class Rating(db.Model):
    __tablename__ = 'ratings'

    id = db.Column(db.String(255), primary_key=True)
    comment = db.Column(db.String(255))
    rating = db.Column(db.Integer, nullable=False)
    rating_date = db.Column(db.String(255), nullable=False)
    song_id = db.Column(db.String(255), db.ForeignKey('songs.Id'), nullable=False)
    user_id = db.Column(db.String(255), nullable=False)

    def __init__(self, comment, rating, rating_date, song_id, user_id):
        self.comment = comment
        self.rating = rating
        self.rating_date = rating_date
        self.song_id = song_id
        self.user_id = user_id

    @staticmethod
    def from_dict(obj):
        return Rating(
            comment=obj.get("comment"),
            rating=obj.get("rating"),
            rating_date=obj.get("rating_date"),
            song_id=obj.get("song_id"),
            user_id=obj.get("user_id")
        )

