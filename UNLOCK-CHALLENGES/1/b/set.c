#include <stdio.h>

// struct - allows you to group multiple variables of different data types together under a single name.
struct Person {
    char name[20];
    int age;
};

int main() {
    struct Person person1;

    // strcpy - is a standard library function that is used to copy a string from one character array (source) to another character array (destination). The name strcpy stands for "string copy".
    strcpy(person1.name, "John");
    person1.age = 30;

    // %s - indicate the location where a string should be printed or read
    // \n - represent a newline or line break in a string.
    printf("Name: %s\n", person1.name);
    printf("Age: %d\n", person1.age);

    return 0;
}
