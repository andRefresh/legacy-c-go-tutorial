cmake_minimum_required(VERSION 3.1...3.30)

project(
    mmf 
    LANGUAGES C
)

set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/lib)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)

add_subdirectory(mmf)
link_libraries(mmf)

add_subdirectory(legacy_c)
