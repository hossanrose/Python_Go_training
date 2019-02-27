grades = [23, 343, 43, 343]  #List is mutable
tuple_grades =  (34, 343 ,12, 23,34) #Tuple is imutable
set_grades = {3, 23, 52, 23, 83} # set elements should be unique and unordered


print(sum(grades)/len(grades))

# List operations
grades.append(200)
print(grades)
print(grades[0])


#set operations
set_grades.add(300)
print(set_grades)
set_another ={3, 23, 35, 29}
print(set_grades.intersection(set_another))
