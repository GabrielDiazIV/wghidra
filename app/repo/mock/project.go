package mock

var (
	Project = `{
	"project_id" : "my-project-id",
  "functions": [
    {
      "body": "#include \"__gmon_start__@00105040.h\"\n\nint _init(EVP_PKEY_CTX *ctx)\n\n{\n  undefined *puVar1;\n  \n  puVar1 = PTR___gmon_start___00103fd0;\n  if (PTR___gmon_start___00103fd0 != (undefined *)0x0) {\n    puVar1 = (undefined *)(*(code *)PTR___gmon_start___00103fd0)();\n  }\n  return (int)puVar1;\n}\n\n",
      "name": "_init@00101000",
      "parameters": [
        "EVP_PKEY_CTX *"
      ]
    },
    {
      "body": "\nvoid FUN_00101020(void)\n\n{\n                    /* WARNING: Treating indirect jump as call */\n  (*(code *)PTR_00103ff8)();\n  return;\n}\n\n",
      "name": "FUN_00101020@00101020",
      "parameters": []
    },
    {
      "body": "#include \"__cxa_atexit@00105018.h\"\n\nvoid __cxa_atexit(void)\n\n{\n  (*(code *)PTR___cxa_atexit_00104000)();\n  return;\n}\n\n",
      "name": "__cxa_atexit@00101030",
      "parameters": []
    },
    {
      "body": "#include \"operator\u003c\u003c@00105020.h\"\n\n/* WARNING: Unknown calling convention -- yet parameter storage is locked */\n\nbasic_ostream * std::operator\u003c\u003c(basic_ostream *param_1,char *param_2)\n\n{\n  basic_ostream *pbVar1;\n  \n  pbVar1 = (basic_ostream *)(*(code *)PTR_operator\u003c\u003c_00104008)();\n  return pbVar1;\n}\n\n",
      "name": "operator\u003c\u003c@00101040",
      "parameters": [
        "basic_ostream *",
        "char *"
      ]
    },
    {
      "body": "#include \"operator\u003c\u003c@00105028.h\"\n\nvoid __thiscall\nstd::basic_ostream\u003cchar,std::char_traits\u003cchar\u003e\u003e::operator\u003c\u003c\n          (basic_ostream\u003cchar,std::char_traits\u003cchar\u003e\u003e *this,\n          _func_basic_ostream_ptr_basic_ostream_ptr *param_1)\n\n{\n  (*(code *)PTR_operator\u003c\u003c_00104010)();\n  return;\n}\n\n",
      "name": "operator\u003c\u003c@00101050",
      "parameters": [
        "basic_ostream\u003cchar,std::char_traits\u003cchar\u003e\u003e *",
        "_func_basic_ostream_ptr_basic_ostream_ptr *"
      ]
    },
    {
      "body": "#include \"Init@00105030.h\"\n\nvoid __thiscall std::ios_base::Init::Init(Init *this)\n\n{\n  (*(code *)PTR_Init_00104018)();\n  return;\n}\n\n",
      "name": "Init@00101060",
      "parameters": [
        "Init *"
      ]
    },
    {
      "body": "#include \"__libc_start_main@00105010.h\"\n#include \"main@001011a0.h\"\n\nvoid _start(undefined8 param_1,undefined8 param_2,undefined8 param_3)\n\n{\n  undefined8 unaff_retaddr;\n  undefined auStack_8 [8];\n  \n  (*(code *)PTR___libc_start_main_00103fc0)\n            (main,unaff_retaddr,\u0026stack0x00000008,0,0,param_3,auStack_8);\n  do {\n                    /* WARNING: Do nothing block with infinite loop */\n  } while( true );\n}\n\n",
      "name": "_start@00101070",
      "parameters": []
    },
    {
      "body": "#include \"_ITM_deregisterTMCloneTable@00105038.h\"\n\n/* WARNING: Removing unreachable block (ram,0x001010b3) */\n/* WARNING: Removing unreachable block (ram,0x001010bf) */\n\nvoid FUN_001010a0(void)\n\n{\n  return;\n}\n\n",
      "name": "FUN_001010a0@001010a0",
      "parameters": []
    },
    {
      "body": "#include \"_ITM_registerTMCloneTable@00105048.h\"\n\n/* WARNING: Removing unreachable block (ram,0x001010f4) */\n/* WARNING: Removing unreachable block (ram,0x00101100) */\n\nvoid FUN_001010d0(void)\n\n{\n  return;\n}\n\n",
      "name": "FUN_001010d0@001010d0",
      "parameters": []
    },
    {
      "body": "#include \"__cxa_finalize@00105000.h\"\n#include \"FUN_001010a0@001010a0.h\"\n\nvoid FUN_00101110(void)\n\n{\n  if (DAT_00104150 == '\\0') {\n    if (PTR___cxa_finalize_00103fb0 != (undefined *)0x0) {\n      (*(code *)PTR___cxa_finalize_00103fb0)(__dso_handle);\n    }\n    FUN_001010a0();\n    DAT_00104150 = 1;\n    return;\n  }\n  return;\n}\n\n",
      "name": "FUN_00101110@00101110",
      "parameters": []
    },
    {
      "body": "#include \"FUN_001010d0@001010d0.h\"\n\nvoid thunk_FUN_001010d0(void)\n\n{\n  FUN_001010d0();\n  return;\n}\n\n",
      "name": "thunk_FUN_001010d0@00101160",
      "parameters": []
    },
    {
      "body": "\n/* add(int, int) */\n\nint add(int param_1,int param_2)\n\n{\n  return param_2 + param_1;\n}\n\n",
      "name": "add@00101169",
      "parameters": [
        "int",
        "int"
      ]
    },
    {
      "body": "#include \"add@00101169.h\"\n\n/* mult(int, int) */\n\nint mult(int param_1,int param_2)\n\n{\n  int iVar1;\n  \n  iVar1 = add(param_1,param_2);\n  return iVar1 * param_2;\n}\n\n",
      "name": "mult@0010117d",
      "parameters": [
        "int",
        "int"
      ]
    },
    {
      "body": "#include \"operator\u003c\u003c@00101040.h\"\n#include \"endl\u003cchar,std::char_traits\u003cchar\u003e\u003e@00105008.h\"\n#include \"operator\u003c\u003c@00101050.h\"\n\nundefined8 main(void)\n\n{\n  basic_ostream *this;\n  \n  this = std::cout \u003c\u003c \"Hello, World!\";\n  std::basic_ostream\u003cchar,std::char_traits\u003cchar\u003e\u003e::operator\u003c\u003c\n            ((basic_ostream\u003cchar,std::char_traits\u003cchar\u003e\u003e *)this,\n             (_func_basic_ostream_ptr_basic_ostream_ptr *)\n             PTR_endl\u003cchar,std_char_traits\u003cchar\u003e\u003e_00103fb8);\n  return 0;\n}\n\n",
      "name": "main@001011a0",
      "parameters": []
    },
    {
      "body": "#include \"~Init@00105050.h\"\n#include \"Init@00101060.h\"\n#include \"__cxa_atexit@00101030.h\"\n\n/* __static_initialization_and_destruction_0(int, int) */\n\nvoid __static_initialization_and_destruction_0(int param_1,int param_2)\n\n{\n  if ((param_1 == 1) \u0026\u0026 (param_2 == 0xffff)) {\n    std::ios_base::Init::Init((Init *)\u0026std::__ioinit);\n    __cxa_atexit(PTR_~Init_00103fe0,\u0026std::__ioinit,\u0026__dso_handle);\n  }\n  return;\n}\n\n",
      "name": "__static_initialization_and_destruction_0@001011d6",
      "parameters": [
        "int",
        "int"
      ]
    },
    {
      "body": "#include \"__static_initialization_and_destruction_0@001011d6.h\"\n\n/* WARNING: Unknown calling convention -- yet parameter storage is locked */\n/* add(int, int) */\n\nvoid _GLOBAL__sub_I_add(void)\n\n{\n  __static_initialization_and_destruction_0(1,0xffff);\n  return;\n}\n\n",
      "name": "_GLOBAL__sub_I_add@00101228",
      "parameters": []
    },
    {
      "body": "\nvoid _fini(void)\n\n{\n  return;\n}\n\n",
      "name": "_fini@00101240",
      "parameters": []
    },
    {
      "body": "\n/* WARNING: Control flow encountered bad instruction data */\n\nvoid __cxa_finalize(void)\n\n{\n                    /* WARNING: Bad instruction - Truncating control flow here */\n                    /* __cxa_finalize@GLIBC_2.2.5 */\n  halt_baddata();\n}\n\n",
      "name": "__cxa_finalize@00105000",
      "parameters": []
    },
    {
      "body": "\n/* WARNING: Control flow encountered bad instruction data */\n/* WARNING: Unknown calling convention -- yet parameter storage is locked */\n\nbasic_ostream * std::endl\u003cchar,std::char_traits\u003cchar\u003e\u003e(basic_ostream *param_1)\n\n{\n                    /* WARNING: Bad instruction - Truncating control flow here */\n                    /* _ZSt4endlIcSt11char_traitsIcEERSt13basic_ostreamIT_T0_ES6_@GLIBCXX_3.4 */\n  halt_baddata();\n}\n\n",
      "name": "endl\u003cchar,std::char_traits\u003cchar\u003e\u003e@00105008",
      "parameters": [
        "basic_ostream *"
      ]
    },
    {
      "body": "\n/* WARNING: Control flow encountered bad instruction data */\n\nvoid __libc_start_main(void)\n\n{\n                    /* WARNING: Bad instruction - Truncating control flow here */\n                    /* __libc_start_main@GLIBC_2.34 */\n  halt_baddata();\n}\n\n",
      "name": "__libc_start_main@00105010",
      "parameters": []
    },
    {
      "body": "\n/* WARNING: Control flow encountered bad instruction data */\n\nvoid __cxa_atexit(void)\n\n{\n                    /* WARNING: Bad instruction - Truncating control flow here */\n                    /* __cxa_atexit@GLIBC_2.2.5 */\n  halt_baddata();\n}\n\n",
      "name": "__cxa_atexit@00105018",
      "parameters": []
    },
    {
      "body": "\n/* WARNING: Control flow encountered bad instruction data */\n/* WARNING: Unknown calling convention -- yet parameter storage is locked */\n\nbasic_ostream * std::operator\u003c\u003c(basic_ostream *param_1,char *param_2)\n\n{\n                    /* WARNING: Bad instruction - Truncating control flow here */\n                    /* _ZStlsISt11char_traitsIcEERSt13basic_ostreamIcT_ES5_PKc@GLIBCXX_3.4 */\n  halt_baddata();\n}\n\n",
      "name": "operator\u003c\u003c@00105020",
      "parameters": [
        "basic_ostream *",
        "char *"
      ]
    },
    {
      "body": "\n/* WARNING: Control flow encountered bad instruction data */\n\nvoid __thiscall\nstd::basic_ostream\u003cchar,std::char_traits\u003cchar\u003e\u003e::operator\u003c\u003c\n          (basic_ostream\u003cchar,std::char_traits\u003cchar\u003e\u003e *this,\n          _func_basic_ostream_ptr_basic_ostream_ptr *param_1)\n\n{\n                    /* WARNING: Bad instruction - Truncating control flow here */\n                    /* _ZNSolsEPFRSoS_E@GLIBCXX_3.4 */\n  halt_baddata();\n}\n\n",
      "name": "operator\u003c\u003c@00105028",
      "parameters": [
        "basic_ostream\u003cchar,std::char_traits\u003cchar\u003e\u003e *",
        "_func_basic_ostream_ptr_basic_ostream_ptr *"
      ]
    },
    {
      "body": "\n/* WARNING: Control flow encountered bad instruction data */\n\nvoid __thiscall std::ios_base::Init::Init(Init *this)\n\n{\n                    /* WARNING: Bad instruction - Truncating control flow here */\n                    /* _ZNSt8ios_base4InitC1Ev@GLIBCXX_3.4 */\n  halt_baddata();\n}\n\n",
      "name": "Init@00105030",
      "parameters": [
        "Init *"
      ]
    },
    {
      "body": "\n/* WARNING: Control flow encountered bad instruction data */\n\nvoid _ITM_deregisterTMCloneTable(void)\n\n{\n                    /* WARNING: Bad instruction - Truncating control flow here */\n  halt_baddata();\n}\n\n",
      "name": "_ITM_deregisterTMCloneTable@00105038",
      "parameters": []
    },
    {
      "body": "\n/* WARNING: Control flow encountered bad instruction data */\n\nvoid __gmon_start__(void)\n\n{\n                    /* WARNING: Bad instruction - Truncating control flow here */\n  halt_baddata();\n}\n\n",
      "name": "__gmon_start__@00105040",
      "parameters": []
    },
    {
      "body": "\n/* WARNING: Control flow encountered bad instruction data */\n\nvoid _ITM_registerTMCloneTable(void)\n\n{\n                    /* WARNING: Bad instruction - Truncating control flow here */\n  halt_baddata();\n}\n\n",
      "name": "_ITM_registerTMCloneTable@00105048",
      "parameters": []
    },
    {
      "body": "\n/* WARNING: Control flow encountered bad instruction data */\n\nvoid __thiscall std::ios_base::Init::~Init(Init *this)\n\n{\n                    /* WARNING: Bad instruction - Truncating control flow here */\n                    /* _ZNSt8ios_base4InitD1Ev@GLIBCXX_3.4 */\n  halt_baddata();\n}\n\n",
      "name": "~Init@00105050",
      "parameters": [
        "Init *"
      ]
    }
  ],
  "assembly": "endbr64\nsub rsp,0x8\nmov rax,qword ptr [0x00103fd0]\ntest rax,rax\njz 0x00101016\ncall rax\nadd rsp,0x8\nret\npush qword ptr [0x00103ff0]\njmp qword ptr [0x00103ff8]\njmp qword ptr [0x00104000]\npush 0x0\njmp 0x00101020\njmp qword ptr [0x00104008]\npush 0x1\njmp 0x00101020\njmp qword ptr [0x00104010]\npush 0x2\njmp 0x00101020\njmp qword ptr [0x00104018]\npush 0x3\njmp 0x00101020\nendbr64\nxor ebp,ebp\nmov r9,rdx\npop rsi\nmov rdx,rsp\nand rsp,-0x10\npush rax\npush rsp\nxor r8d,r8d\nxor ecx,ecx\nlea rdi,[0x1011a0]\ncall qword ptr [0x00103fc0]\nhlt\nlea rdi,[0x104030]\nlea rax,[0x104030]\ncmp rax,rdi\njz 0x001010c8\nmov rax,qword ptr [0x00103fc8]\ntest rax,rax\njz 0x001010c8\njmp rax\nret\nlea rdi,[0x104030]\nlea rsi,[0x104030]\nsub rsi,rdi\nmov rax,rsi\nshr rsi,0x3f\nsar rax,0x3\nadd rsi,rax\nsar rsi,0x1\njz 0x00101108\nmov rax,qword ptr [0x00103fd8]\ntest rax,rax\njz 0x00101108\njmp rax\nret\nendbr64\ncmp byte ptr [0x00104150],0x0\njnz 0x00101150\npush rbp\ncmp qword ptr [0x00103fb0],0x0\nmov rbp,rsp\njz 0x00101138\nmov rdi,qword ptr [0x00104028]\ncall qword ptr [0x00103fb0]\ncall 0x001010a0\nmov byte ptr [0x00104150],0x1\npop rbp\nret\nret\nendbr64\njmp 0x001010d0\npush rbp\nmov rbp,rsp\nmov dword ptr [rbp + -0x4],edi\nmov dword ptr [rbp + -0x8],esi\nmov edx,dword ptr [rbp + -0x4]\nmov eax,dword ptr [rbp + -0x8]\nadd eax,edx\npop rbp\nret\npush rbp\nmov rbp,rsp\nsub rsp,0x8\nmov dword ptr [rbp + -0x4],edi\nmov dword ptr [rbp + -0x8],esi\nmov edx,dword ptr [rbp + -0x8]\nmov eax,dword ptr [rbp + -0x4]\nmov esi,edx\nmov edi,eax\ncall 0x00101169\nimul eax,dword ptr [rbp + -0x8]\nleave\nret\npush rbp\nmov rbp,rsp\nlea rax,[0x102004]\nmov rsi,rax\nlea rax,[0x104040]\nmov rdi,rax\ncall 0x00101040\nmov rdx,qword ptr [0x00103fb8]\nmov rsi,rdx\nmov rdi,rax\ncall 0x00101050\nmov eax,0x0\npop rbp\nret\npush rbp\nmov rbp,rsp\nsub rsp,0x10\nmov dword ptr [rbp + -0x4],edi\nmov dword ptr [rbp + -0x8],esi\ncmp dword ptr [rbp + -0x4],0x1\njnz 0x00101225\ncmp dword ptr [rbp + -0x8],0xffff\njnz 0x00101225\nlea rax,[0x104151]\nmov rdi,rax\ncall 0x00101060\nlea rax,[0x104028]\nmov rdx,rax\nlea rax,[0x104151]\nmov rsi,rax\nmov rax,qword ptr [0x00103fe0]\nmov rdi,rax\ncall 0x00101030\nnop\nleave\nret\npush rbp\nmov rbp,rsp\nmov esi,0xffff\nmov edi,0x1\ncall 0x001011d6\npop rbp\nret\nendbr64\nsub rsp,0x8\nadd rsp,0x8\nret\n"
}`
)
