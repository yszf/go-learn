#include "./person/person2.h"

#include <stdio.h>

// go build -buildmode=c-archive -o test_person.a test_person.go
// g++ test_person3.cc test_person.a -lpthread -o test_person3 -std=c++11
int main() {
    auto p = Person::New("mary", 15);

    char buf[64];
    char* name = p->GetName(buf, sizeof(buf)-1);
    int age = p->GetAge();

    printf("%s, %d years old.\n", name, age);

    p->Delete();

    return 0;
}