### Dependencies
1. `/joho/godotenv`
2. `/CezarGarrido/sqllogs`
3. `/labstack/gommon/log`
4. `/gofiber/fiber/v2`
5. `"/dgrijalva/jwt-go`



5. `/go-sql-driver/mysql`

### Abbreviation and Convention
1. > **iiod** is _Init, Interface, Override, Dependency_
2. > ~~**port** is _Port in Application Running_~~
3. > ~~**s** is _Variable Struct_~~
4. > // **ValidateRequestAdmin TODO** satu file yang dipakai ditempat kain

### Step by Step
1. dalam override MEMPUNYAI bertipe type data interface lain
2. dalam override terdapat task memanggil INIT interface lain

registry akan mendaftarkannya dalam satu modul, crud admin,crud apapun

repo untuk sementara langsung mysql, 
nanti bila banyak perlu ganti mysql, mongo

> runtime error: invalid memory address or nil pointer dereference
disebabkan response belum di instance

common, -> constant, helper, 