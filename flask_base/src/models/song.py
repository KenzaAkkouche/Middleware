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


