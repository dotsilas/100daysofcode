int[] a = {1,2,3,4,5};

Console.WriteLine("Contents of Array:");
PrintArray();

void PrintArray()
{
    foreach (int x in a)
    {
        Console.Write($"{x} ");
    }
    Console.WriteLine();
}

// Method to create random numbers

void DisplayRandomNumbers()
{
    Random random = new Random();
    
    for (int i = 0; i < 5; i++)
    {
        Console.Write($"{random.Next(1,100)} ");
    }
    
    Console.WriteLine();
}

Console.WriteLine("Generating random numbers");
DisplayRandomNumbers();

// Parameters passed by value or reference

// Arrays are reference types
int[] array = {1, 2, 3, 4, 5};

PrintArray2(array);
Clear(array);
PrintArray2(array);

void PrintArray2(int[] array) 
{
    foreach (int a in array) 
    {
        Console.Write($"{a} ");
    }
    Console.WriteLine();
}

void Clear(int[] array) 
{
    for (int i = 0; i < array.Length; i++) 
    {
        array[i] = 0;
    }
}

// Named Parameters
// PrintPersonalData(firstname: "John", age: 25, lastname: "Marx");

// Optional parameters has default values assigned
// int Sum(firstNumber, secondNumber = 0) { return firstname + secondNumber; }
// int result = Sum(5);
// result == 5


