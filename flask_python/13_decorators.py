import functools

def my_decorator(func):
    functools.wraps(func)  # This line makes it a my_decorator
    def function_that_runs_func():
        print ("Before")
        func()
        print("after")
    return function_that_runs_func


### The below fuction will get replaced by the decorator fucntion

@my_decorator   ##Adding decorator makes this function a decorator
def my_function():     # My funcion gets replaced by
    print ("Middle")

my_function()

##### Decorator with arguments

def decorator_with_args(number):
    def my_decorator(func):
        functools.wraps(func)
        def function_that_runs_func(*args, **kwargs):
            print("Before")
            if number == 24:    #Using the decorator argument
                func(*args, **kwargs)
            else:
                print ("test")
            print("After")
        return function_that_runs_func
    return my_decorator


@decorator_with_args(24)   #Decorator with argument
def my_function_2(x, y):
    print (x+y)

my_function_2(5,6)
