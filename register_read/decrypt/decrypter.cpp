#include <Windows.h>
#include <iostream>
#include <memory>
#include <vector>
#include <algorithm>
#include <fstream>

// try to get cpp decrypter as class function
class CppDecrypter{
	public:
		CppDecrypter(); //Konstruktor
		~CppDecrypter(); //Destruktor
		string decrypt(string sk); // Decrypter function
	
	private:
		auto get_decryptfunc();
		constexpr auto CIPHER_SIZE;
};

CppDecrypter::CppDecrypter()
: CIPHER_SIZE(0x80)
{}

CppDecrypter::~CppDecrypter(){}

// constexpr auto CIPHER_SIZE = 0x80;

void fail(std::string&& msg) 
{
	std::cerr << msg;
	std::cin.get();
	exit(1);
}

const auto CppDecrypter::get_decryptfunc()
{
	constexpr auto func_dec_offset = 0x85249;

	const auto ipcsecproc = reinterpret_cast<char*>(LoadLibraryA("ipcsecproc.dll"));

	if (ipcsecproc == nullptr)
	{
		fail("Couldn't load ipcsecproc.dll, make sure it's loadeable!\n");
	}

	return reinterpret_cast<bool(__stdcall* const)(char * dec, const char * enc)>(ipcsecproc + func_dec_offset);
}


string CppDecrypter::decrypt(string sk)
{

	const auto func_dec = get_decryptfunc();
	const auto cipher = sk;
	std::vector<char> plain(0x80);

	if (func_dec(plain.data(), cipher)
	{
		std::reverse(plain.begin(), plain.begin() + 0x30);
		return plain
	}

	return "Error";
}

// auto CppDecrypter::get_cipher(std::string&& filename)
// {
// 	std::ifstream ifs{ filename, std::ios::binary | std::ios::ate };
// 	if (ifs.is_open())
// 	{
// 		const auto size = ifs.tellg(); ifs.seekg(0);
// 		if (size < CIPHER_SIZE)
// 		{
// 			fail("File ins't big enough for decryption!\n");
// 		}
		
// 		auto cipher = std::vector<char>(std::istreambuf_iterator<char>{ ifs }, {});
// 		std::reverse(cipher.begin(), cipher.end());
// 		return cipher;
// 	}
// 	else
// 	{
// 		fail("Couldn't open cipher file, make sure it's readable!\n");
// 	}
// }

// void write_plain(std::string&& filename, const std::vector<char>& plain)
// {
// 	std::ofstream ofs(filename, std::ios::binary | std::ios::trunc);
// 	if (ofs.is_open()) {
// 		ofs.write(plain.data(), plain.size());
// 	}
// 	else
// 	{
// 		fail("Couldn't open plaintext file, make sure it's writeable!\n");
// 	}
// }
