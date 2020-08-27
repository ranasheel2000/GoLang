start=int(input("Enter the interval start value"))
end=int(input("Enter the interval end value"))

for i in range(start,end+1):
    l=len(str(i))
    sum=0
    tar=i
    while(i>0):
        rem=i%10
        i=i//10
        sum=sum+rem**l
    if(sum==tar):
        print("Number={0} is armstrong".format(tar))
        
