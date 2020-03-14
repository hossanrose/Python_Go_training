details = {
    'name' : "Hossan",
    'age' : 30,
    'grades' : [10,23, 45, 83]
}

#Print out a value using a key
print (details["grades"])

#iterating through the dictionary
for key, value in details.items():
    print(f"{key}:{value}")
