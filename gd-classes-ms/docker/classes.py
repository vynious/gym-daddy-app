from flask import Flask, request, jsonify
from flask_cors import CORS
from flask_sqlalchemy import SQLAlchemy
from datetime import datetime
from os import environ
from sqlalchemy import func


app = Flask(__name__)
CORS(app)
app.config['SQLALCHEMY_DATABASE_URI'] = environ.get('dbURL')
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False


db = SQLAlchemy(app)


class Classes(db.Model):
    __tablename__ = 'classes'


    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    name = db.Column(db.String(255), nullable=False)
    description = db.Column(db.Text, nullable=False)
    duration = db.Column(db.Integer)
    date_time = db.Column(db.DateTime)
    suitable_level = db.Column(db.String(255))
    capacity = db.Column(db.Integer)
    max_capacity = db.Column(db.Integer, nullable=False)

    def __init__(self, name, description, duration, date_time, suitable_level, capacity, max_capacity):
        self.name = name
        self.description = description
        self.duration = duration
        self.date_time = date_time
        self.suitable_level = suitable_level
        self.capacity = capacity
        self.max_capacity = max_capacity


    def json(self):
        return {"id": self.id,
                "name": self.name,
                "description": self.description,
                "duration": self.duration,
                "date_time": self.date_time,
                "suitable_level": self.suitable_level,
                "capacity": self.capacity,
                'max_capacity': self.max_capacity}



@app.route("/api/classes", methods=['GET', 'POST'])
def classes():
    if request.method == 'GET':
        class_list = Classes.query.all()
        classes_data = [class_item.json() for class_item in class_list]
        return jsonify({"classes": classes_data})

    elif request.method == 'POST':
        data = request.get_json()

        new_class = Classes(
            name=data.get('name'),
            description=data.get('description'),
            duration=data.get('duration'),
            date_time=datetime.strptime(data.get('date_time'), '%Y-%m-%d %H:%M:%S'),
            suitable_level=data.get('suitable_level'),
            capacity=data.get('capacity'),
            max_capacity=data.get('capacity')
        )
        try:
            db.session.add(new_class)
            db.session.commit()
            return jsonify({"message": "Class created successfully", "class": new_class.json()})
        except Exception as e:
            return jsonify({"message": "Failed to create class", "error": str(e)}), 500



@app.route("/api/classes/<int:id>", methods=['GET', 'DELETE', 'PATCH'])
def manage_class(id):
    if request.method == 'GET':
        class_item = Classes.query.get(id)
        if class_item:
            return jsonify(
                {
                    "code": 200,
                    "data": class_item.json()
                }
            ), 200
        return jsonify(
            {
                "code": 404,
                "message": "Class not found."
            }
        ), 404
        
    elif request.method == 'DELETE':
        class_item = Classes.query.get(id)

        if class_item:
            db.session.delete(class_item)
            db.session.commit()
            return jsonify(
                {
                    "code": 200,
                    "message": "Class deleted successfully."
                    }
                ), 200
        else:
            return jsonify(
                {
                    "code": 404,
                    "message": "Class not found."
                }
            ), 404

    elif request.method == 'PATCH':   
        try:
            data = request.get_json()
            action = data.get('action')

            if action not in ['book', 'cancel']:
                return jsonify(
                    {
                        "code": 400,
                        "message": "Invalid action. Please choose either 'book' or 'cancel'."
                    }
                ), 400

            class_item = Classes.query.get(id)

            if class_item:
                if action == 'book':
                    if class_item.capacity <= 0:
                        return jsonify(
                            {
                                "code": 400,
                                "message": "Class is fully booked."
                            }
                        ), 400
                    class_item.capacity -= 1
                elif action == 'cancel':
                    if class_item.capacity >= class_item.max_capacity:
                        return jsonify({"code": 400, "message": "Class is already at maximum capacity."}), 400
                    class_item.capacity += 1

                db.session.commit()
                return jsonify(
                    {
                        "code": 201,
                        "message": "Class capacity updated successfully."
                    }
                ), 201
            else:
                return jsonify(
                    {
                        "code": 404,
                        "message": "Class not found."
                    }
                ), 404
        except Exception as e:
            print(str(e))
            return jsonify(
                {
                    "code": 500,
                    "message": f"An error occurred: {str(e)}"
                }
            ), 500

    

@app.route("/api/classes/date/<string:date>", methods=['GET'])
def get_schedule_for_specific_day(date):
    try:
        target_date = datetime.strptime(date, '%Y-%m-%d').date()

        classList = Classes.query.filter(func.DATE(Classes.date_time) == target_date).all()

        classes_data = []
        for class_item in classList:
            date_time_str = class_item.date_time.strftime('%Y-%m-%d %H:%M:%S')

            class_info = {
                "id": class_item.id,
                "name": class_item.name,
                "description": class_item.description,
                "duration": class_item.duration,
                "date_time": date_time_str,
                "suitable_level": class_item.suitable_level,
                "capacity": class_item.capacity,
                "max_capacity": class_item.max_capacity
            }
            classes_data.append(class_info)

        if len(classes_data):
            return jsonify(
                {
                    "code": 200,
                    "data": {"classes": classes_data}
                }
            )
        else:
            return jsonify(
                {
                    "code": 404,
                    "message": "There are no classes scheduled for the specified date."
                }
            ), 404
    except ValueError:
        return jsonify(
            {
                "code": 400,
                "message": "Invalid date format. Please use the format YYYY-MM-DD."
            }
        ), 400
    except Exception as e:
        
        return jsonify(
            {
                "code": 500,
                "message": f"An error occurred: {str(e)}"
            }
        ), 500




if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5200, debug=True)