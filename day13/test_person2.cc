#include "./person/person2.h"

#include <stdio.h>

// go build -buildmode=c-archive -o person_capi.a person_capi.go
// g++ test_person2.cc person_capi.a -lpthread -o test_person2 -std=c++11
int main() {
    auto p = Person::New("jack", 20);

    char buf[64];
    char* name = p->GetName(buf, sizeof(buf)-1);
    int age = p->GetAge();

    printf("%s, %d years old.\n", name, age);

    p->Delete();

    return 0;
}