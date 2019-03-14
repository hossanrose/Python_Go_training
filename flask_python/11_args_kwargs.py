"""*args can be used to pass in any number of arguments as a tulip"""
def my_args(*args):
    print(args)

"""**kwargs can be used to pass in named arguments as a set """
def my_kwargs(**kwargs):
    print(kwargs)

"""args and kwargs can be combined"""

def my_combined(*args,**kwargs):
    print(args,kwargs)

my_args(1,2,3,4)
my_kwargs(name="Test",Yo="Testing")
my_combined(1,2,3,4,name="Test",Yo="Testing")
