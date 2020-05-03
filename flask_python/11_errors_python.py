def divide (dividend, divisor):
    if divisor == 0:
        raise ZeroDivisionError("Divisor cannot be 0")
    return dividend/divisor

grades =[]

print("Welcome to average grade program")
try:                                # Try block is used to catch the raised exception
    average = divide(sum(grades),len(grades))
except ZeroDivisionError as e:      # e has the value of the error
    print("There are no grades yet in your list")
else:
    print(f"The average grade is {average}")
