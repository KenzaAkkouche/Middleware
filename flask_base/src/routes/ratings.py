from flask import Blueprint, jsonify, request
import requests
from src.models.song import Song
import src.routes.songs as song_routes

ratings_app = Blueprint('ratings_app', __name__)

BASE_URL = 'https://ratings-charlie.edu.forestier.re'

# Routes for operations on ratings

@ratings_app.route('/songs/<song_id>/ratings', methods=['GET'])
def get_song_ratings(song_id):
    ratings_url = f'{BASE_URL}/songs/{song_id}/ratings'

    song = song_routes.get_song(song_id)
    if song is None:
        return jsonify({'error': 'Song not found'}), 404
    try:
        response = requests.get(ratings_url)
        if response.status_code == 200:
            ratings_data = response.json()
            return jsonify(ratings_data)
        else:
            return jsonify({'error': 'Failed to fetch ratings'}), response.status_code

    except requests.RequestException as e:
        return jsonify({'error': f'Request Exception: {str(e)}'}), 500

@ratings_app.route('/songs/<song_id>/ratings', methods=['POST'])
def add_song_rating(song_id):
    ratings_url = f'{BASE_URL}/songs/{song_id}/ratings'

    try:
        data = request.json

        response = requests.post(ratings_url, json=data)
        if response.status_code == 201:  # 201 for Created
            new_rating = response.json()
            return jsonify(new_rating), 201
        else:
            return jsonify({'error': 'Failed to add rating'}), response.status_code

    except requests.RequestException as e:
        return jsonify({'error': f'Request Exception: {str(e)}'}), 500

@ratings_app.route('/songs/<song_id>/ratings/<rating_id>', methods=['GET'])
def get_specific_rating(song_id, rating_id):
    rating_url = f'{BASE_URL}/songs/{song_id}/ratings/{rating_id}'

    try:
        response = requests.get(rating_url)
        if response.status_code == 200:
            rating_data = response.json()
            return jsonify(rating_data)
        else:
            return jsonify({'error': 'Failed to fetch rating'}), response.status_code

    except requests.RequestException as e:
        return jsonify({'error': f'Request Exception: {str(e)}'}), 500

@ratings_app.route('/songs/<song_id>/ratings/<rating_id>', methods=['PUT'])
def update_rating(song_id, rating_id):
    rating_url = f'{BASE_URL}/songs/{song_id}/ratings/{rating_id}'

    try:
        data = request.json

        response = requests.put(rating_url, json=data)
        if response.status_code == 200:
            updated_rating = response.json()
            return jsonify(updated_rating)
        else:
            return jsonify({'error': 'Failed to update rating'}), response.status_code

    except requests.RequestException as e:
        return jsonify({'error': f'Request Exception: {str(e)}'}), 500

@ratings_app.route('/songs/<song_id>/ratings/<rating_id>', methods=['DELETE'])
def delete_rating(song_id, rating_id):
    rating_url = f'{BASE_URL}/songs/{song_id}/ratings/{rating_id}'

    try:
        response = requests.delete(rating_url)
        if response.status_code == 204:  # 204 for No Content
            return '', 204
        else:
            return jsonify({'error': 'Failed to delete rating'}), response.status_code

    except requests.RequestException as e:
        return jsonify({'error': f'Request Exception: {str(e)}'}), 500
