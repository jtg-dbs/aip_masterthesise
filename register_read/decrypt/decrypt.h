#ifdef __cplusplus
extern "C" {
#endif
    typedef void* CppDecrypter;
    CppDecrypter CppDecrypterInit(void);
    void CppDecrypterFree(CppDecrypter);
    char* CppDecrypterDecrypt(char str[]);
#ifdef __cplusplus
}
#endif