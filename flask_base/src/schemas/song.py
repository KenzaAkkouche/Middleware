from marshmallow import Schema, fields, validates_schema, ValidationError

class RatingSchema(Schema):
    comment = fields.String(description="Comment")
    id = fields.String(description="ID")
    rating = fields.Integer(description="Rating")
    rating_date = fields.DateTime(description="Rating Date")
    song_id = fields.String(description="Song ID")
    user_id = fields.String(description="User ID")

class SongSchema(Schema):
    Id = fields.String(description="UUID")
    Title = fields.String(description="Title of Song")
    Artist = fields.String(description="Name of Artist")
    ratings = fields.Nested(RatingSchema, many=True)


    @staticmethod
    def is_empty(obj):
        return (not obj.get("Id") or obj.get("Id") == "") and \
               (not obj.get("Title") or obj.get("Title") == "") and \
               (not obj.get("Artist") or obj.get("Artist") == "")

class BaseSongSchema(Schema):
    Title = fields.String(description="Title of song", required=True)
    Artist = fields.String(description="Name of Artist ", required=True)
    ratings = fields.Nested(RatingSchema, many=True)


class SongUpdateSchema(BaseSongSchema):
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("Title" in data and data["Title"] != "") or
                ("Artist" in data and data["Artist"] != "") or
                ("ratings" in data and data["ratings"])):
            raise ValidationError("At least one of ['Title', 'Artist', 'ratings'] must be specified")
