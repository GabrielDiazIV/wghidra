#include <iostream>

int add(int a, int b) { return a + b; }
int mult(int a, int b) { return add(a, b) * b; }

int main() {
    std::cout << "Hello, World!" << std::endl;
    return 0;
}
