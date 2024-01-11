from marshmallow import Schema, fields, validates_schema, ValidationError

class SongSchema(Schema):
    Id = fields.String(description="UUID")
    Title = fields.String(description="Title of Song")
    Artist = fields.String(description="Name of Artist")

    @staticmethod
    def is_empty(obj):
        return (not obj.get("Id") or obj.get("Id") == "") and \
               (not obj.get("Title") or obj.get("Title") == "") and \
               (not obj.get("Artist") or obj.get("Artist") == "")

class BaseSongSchema(Schema):
    Title = fields.String(description="Title of song", required=True)
    Artist = fields.String(description="Name of Artist ", required=True)

class SongUpdateSchema(BaseSongSchema):
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("Title" in data and data["Title"] != "") or
                ("Artist" in data and data["Artist"] != "")):
            raise ValidationError("At least one of ['Title', 'Artist'] must be specified")
