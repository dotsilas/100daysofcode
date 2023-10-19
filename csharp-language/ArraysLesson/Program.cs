string[] fraudulentOrderIDs = new string[3];

fraudulentOrderIDs[0] = "A123";
fraudulentOrderIDs[1] = "B456";
fraudulentOrderIDs[2] = "C789";

Console.WriteLine($"First: {fraudulentOrderIDs[0]}");
Console.WriteLine($"Second: {fraudulentOrderIDs[1]}");
Console.WriteLine($"Third: {fraudulentOrderIDs[2]}");
Console.Write("\n");

fraudulentOrderIDs[0] = "F000";

Console.WriteLine($"Reassign First: {fraudulentOrderIDs[0]}");
Console.Write("\n");

// declaration and initialzation of an array

string[] myStrings = {"first", "second", "third"};

foreach (string str in myStrings) {
    Console.WriteLine(str);
}

Console.WriteLine($"The length of myStrings is {myStrings.Length}");
