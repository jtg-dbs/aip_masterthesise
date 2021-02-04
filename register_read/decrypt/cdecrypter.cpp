#include "decrypt.h"
#include "decrypter.cpp"

CppDecrypter CppDecrypterInit(){
    CppDecrypter * ret = new CppDecrypter();
    return (void*) ret;
}

void CppDecrypterFree(CppDecrypter obj){
    CppDecrypter * cd = (CppDecrypter*) obj;
    delte cd;
}

char* CppDecrypterDecrypt(char str[] sk, CppDecrypter decrypt){
    CppDecrypter * decyrpter = (CppDecrypter*) decrypt;
    return (void*) decrypter -> decrypt(sk);
}