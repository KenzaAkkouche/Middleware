from marshmallow import Schema, fields, validates_schema, ValidationError

class RatingSchema(Schema):
    comment = fields.String(description="Comment")
    id = fields.String(description="id")
    rating = fields.Integer(description="Rating")
    rating_date = fields.DateTime(description="Rating Date")
    song_id = fields.String(description="Song id")
    user_id = fields.String(description="User id")

class SongSchema(Schema):
    id = fields.String(description="UUid")
    title = fields.String(description="title of Song")
    artist = fields.String(description="Name of artist")
    ratings = fields.Nested(RatingSchema, many=True)


    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("title") or obj.get("title") == "") and \
               (not obj.get("artist") or obj.get("artist") == "")

class BaseSongSchema(Schema):
    title = fields.String(description="title of song", required=True)
    artist = fields.String(description="Name of artist ", required=True)
    ratings = fields.Nested(RatingSchema, many=True)


class SongUpdateSchema(BaseSongSchema):
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("title" in data and data["title"] != "") or
                ("artist" in data and data["artist"] != "") or
                ("ratings" in data and data["ratings"])):
            raise ValidationError("At least one of ['title', 'artist', 'ratings'] must be specified")
