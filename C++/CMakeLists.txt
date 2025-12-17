cmake_minimum_required(VERSION 3.16)
project(DataStructuresLab)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Coverage флаги
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fprofile-arcs -ftest-coverage")
set(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} -fprofile-arcs -ftest-coverage")

# Основная программа (демо)
add_executable(ds_demo
    main.cpp
    stack.cpp
    array.cpp
    singly_linked_list.cpp
    doubly_linked_list.cpp
    binary_tree.cpp
    hash_table.cpp
    serialization.cpp
)

# Тесты
add_executable(ds_tests
    tests.cpp
    stack.cpp
    array.cpp
    singly_linked_list.cpp
    doubly_linked_list.cpp
    binary_tree.cpp
    hash_table.cpp
    serialization.cpp
)

# Подключение Google Test
target_include_directories(ds_tests PRIVATE /opt/homebrew/include)
target_link_directories(ds_tests PRIVATE /opt/homebrew/lib)
target_link_libraries(ds_tests gtest gtest_main pthread)

# Настройки компиляции
target_compile_options(ds_demo PRIVATE -Wall -Wextra -O0 -g)
target_compile_options(ds_tests PRIVATE -Wall -Wextra -O0 -g)
