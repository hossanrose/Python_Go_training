class MyClass:
    def __init__(self,list):
        self.name = "Test"
        self.list = list
        
    """Normal instance method of a class, this method can modify objects state
    using passed self parameter that points to object"""
    def print_out(self):
        print ("Myname is {} and I'm a list with {}".format(self.name,self.list))

    """Class method can be used to modify the class state
    using the passed cls paramerter that points to class
    Here we are using the classmethod as a constructor of new class"""
    @classmethod
    def create(cls):
        object = cls([1,2])
        return object

    """Static method doesnt have any information about class or pbject
    Used as a independent method without self or cls parameter"""
    @staticmethod
    def print_list():
        print ("I'm inside a Class")

# Checking instance method by creating a object with the class
obj1 = MyClass([3,4])
obj1.print_out()

# Using classmethod to create a object
obj2 = MyClass.create()
obj2.print_out()

## Using staticmethod
MyClass.print_list()
