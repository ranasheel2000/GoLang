class CustomIterater:
    def __init__(self, count=0):
        self.count = count
    def __iter__(self):
        self.n = 0
        return self
    def __next__(self):
        #print("count={0}".format(self.count))
        if self.n <= self.count:
            result = 2 ** self.n
            self.n += 1
            return result
        else:
            raise StopIteration

# Enter how many power of 2 you wants to print, i.e. 5
numbers = CustomIterater(5)
i = iter(numbers)
j=0
while(True):
    try:
        j=next(i)
        print(j)
    except StopIteration:
        print("end of loop")
        break

