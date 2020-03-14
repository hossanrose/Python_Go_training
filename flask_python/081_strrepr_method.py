
class Student:

    def __init__(self,name,age):
        self.name=name
        self.age=age

    def __str__(self): ##Method used to pass a string back incase the object is printed
        return f"{self.name}: {self.age}"

    def __repr__(self): ##Method used to provide debugging output
        return f"{self.name}: {self.age}"


person=Student("Yo","23")
print(person)
