#include <random>
#include <iostream>

using namespace std;

int main() {
mersenne_twister_engine<unsigned long,32ul,624ul,397ul,31ul,2567483615ul,11ul,4294967295ul,7ul,2636928640ul,15ul,4022730752ul,18ul,1812433253ul> r;

    // got from the first characters of the flag
    r.seed(2003266937);
    cout << "[";
    for(int i = 0; i < 125; i++) {
        cout << "[";
        for(int j = 0; j < 16; j++) {
            cout << ((r() & 0xff)) << (j != 15 ? ",": "");
        }
        cout << "]" << (i != 124 ? "," : "")<< endl;
    }
    cout << "]";
}
