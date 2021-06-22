from flask import Flask,redirect,render_template,jsonify,request, url_for
from write_csv import add_data

app = Flask(__name__)

@app.route("/form", methods = ["POST", "GET"])
def takedata():
    if request.method == "POST":
        student = request.form.get("name")
        addmission = request.form.get("addID")
        branch = request.form.get("branch")
        status = add_data(student, addmission, branch)
        return status
    else:
        return render_template("login.html")


if __name__ == "__main__":
    app.run(debug=True)