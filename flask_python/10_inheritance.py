### Super class
class Student:
    def __init__(self, name, school):
        self.name = name
        self.school = school
        self.mark = []

    def average(self):
        return sum(self.mark) / len(self.mark)

    """ Using the classmethod allows us to use the below method according the
     class which calls it """
    @classmethod
    def friend(cls, origin, friend_name, salary):
        return cls(friend_name, origin.school, salary)  ## Passing object as origin

class WorkingStudent(Student):
    def __init__(self, name, school, salary):
        """ Below line of code initializes the variables name and school
        from Student class"""
        super().__init__(name,school)
        self.salary = salary


anna = Student("Anna","MIT")
print(anna.name)

""""WorkingStudent.friend initializes the method friend as a WorkingStudent
method. According the class used to call the method the method will become
that class method """
friend = WorkingStudent.friend(anna, "Greg", 35)
print(friend.name)
print(friend.salary)
