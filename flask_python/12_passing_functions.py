def methodpassing(another):
    return another()

def add_two():
    return 1+2

print(methodpassing(add_two)) # Passing one function to another 

print(methodpassing(lambda: 1+2)) ## Passng lambda funcion
