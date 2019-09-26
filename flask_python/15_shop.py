from flask import Flask, jsonify, request, render_template

app = Flask (__name__)

stores = [
    {
        "name" : "My store",
        "items" :[
            {
            "name" : 'book',
            "price": 10
            }
        ]
    }
]

@app.route("/")
def home():
    return render_template("index.html")

# Create a shop using post
@app.route("/shop", methods=["POST"])
def create_shop():
    request_data = request.get_json()
    new_store = {
        "name" : request_data["name"],
        "item" : []
    }
    stores.append(new_store)
    return jsonify(new_store)


# Give details of a /shop/item_name with GET
@app.route("/shop/<string:name>")   #String : name is a varibale
def get_shop(name):
    #Iterate over stores and if it finds the store return it
    for dic in stores:
        if ( dic["name"] == name ):
            return jsonify(dic)    # dic is already a dictionary
    return jsonify({"message": 'store not found'} )  # outputing the message in dictionary format


# Give details of a /shop/ with GET
@app.route("/shop/")   #String : name is a varibale
def get_total_shop():
    return jsonify({"stores": stores}) ### jsonify is used to change the stores variable to json format

# Create a shop item using post
@app.route("/shop/<string:name>/item", methods=["POST"])
def create_shop_item(name):
    request_item = request.get_json()
    for dic in stores:
        if (dic["name"] == name):
            new_item = {
            "name" : request_item["name"],
            "price" : request_item["price"]
            }
            dic["items"].append(new_item)
            return jsonify(new_item)
    return jsonify({"message": 'store not found'} )


#  Give shop item using get
@app.route("/shop/<string:name>/item")
def get_shop_item(name):
    #Iterate over stores and if it finds the store return it
    for dic in stores:
        if ( dic["name"] == name ):
            return jsonify({"items": dic['items']})    # outputing items in dictionary format
    return jsonify({"message": 'store not found'} )  # outputing the message in dictionary format

app.run(port=1000)
