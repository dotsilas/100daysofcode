// INTEGER TYPES

// MinValue and MaxValue properties
Console.WriteLine("The range of values of integer types:");

Console.WriteLine($"sbyte\t: {sbyte.MinValue} to {sbyte.MaxValue}");
Console.WriteLine($"short\t: {short.MinValue} to {short.MaxValue}");
Console.WriteLine($"int\t: {int.MinValue} to {int.MaxValue}");
Console.WriteLine($"long\t: {long.MinValue} to {long.MaxValue}");

Console.WriteLine("The range of values of unsigned integer types:");

Console.WriteLine($"byte\t: {byte.MinValue} to {byte.MaxValue}");
Console.WriteLine($"ushort\t: {ushort.MinValue} to {ushort.MaxValue}");
Console.WriteLine($"uint\t: {uint.MinValue} to {uint.MaxValue}");
Console.WriteLine($"ulong\t: {ulong.MinValue} to {ulong.MaxValue}");

Console.WriteLine("Ther Range of Floating point types:");
Console.WriteLine($"float  : {float.MinValue} to {float.MaxValue} (with ~6-9 digits of precision)");
Console.WriteLine($"double : {double.MinValue} to {double.MaxValue} (with ~15-17 digits of precision)");
Console.WriteLine($"decimal: {decimal.MinValue} to {decimal.MaxValue} (with 28-29 digits of precision)");


// REFERENCE TYPES
int[] data = new int[3];

// Casting
float myFloat = 3.1416f;
int myInt = 3;
int result = (int)myFloat + myInt;

Console.WriteLine($"Result 3.1416 (casted to int) + 3 = {result}");

// int to string
string myStringFromInt = myInt.ToString();
Console.WriteLine($"Int to String {myStringFromInt}");

// string to int
int MyIntFromString = int.Parse(myStringFromInt);
Console.WriteLine($"String to int {MyIntFromString}");

// Sort and Reverse
string[] pallets = { "B14", "A11", "B12", "A13" };

Console.WriteLine("Sorted...");
Array.Sort(pallets);
foreach (var pallet in pallets)
{
    Console.WriteLine($"-- {pallet}");
}

Console.WriteLine("");
Console.WriteLine("Reversed...");
Array.Reverse(pallets);
foreach (var pallet in pallets)
{
    Console.WriteLine($"-- {pallet}");
}


// Resize
string[] palletsR = { "B14", "A11", "B12", "A13" };
Console.WriteLine("");

Array.Clear(palletsR, 0, 2);
Console.WriteLine($"Clearing 2 ... count: {palletsR.Length}");
foreach (var pallet in palletsR)
{
    Console.WriteLine($"-- {pallet}");
}

Console.WriteLine("");
// pass a variable by reference **
Array.Resize(ref pallets, 6);
Console.WriteLine($"Resizing 6 ... count: {pallets.Length}");

pallets[4] = "C01";
pallets[5] = "C02";

foreach (var pallet in pallets)
{
    Console.WriteLine($"-- {pallet}");
}


// string to chars
string value = "abc123";
char[] valueArray = value.ToCharArray();
Array.Reverse(valueArray);
// string result = new string(valueArray);
string result2 = String.Join(",", valueArray);
Console.WriteLine(result2);
