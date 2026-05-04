section .text
    global _start

_start:
        jmp short jump_to_call

server:
    pop r12

    push 41
    pop rax
    push 2
    pop rdi
    push 1
    pop rsi
    xor edx, edx
    syscall
    xchg edi, eax

    push 0
    mov eax, 0x901F0002
    push rax
    mov rsi, rsp
    push 16
    pop rdx
    push 49
    pop rax
    syscall

    push 50
    pop rax
    xor esi, esi
    syscall

loop:
    push 43
    pop rax
    xor esi, esi
    xor edx, edx
    syscall
    mov r8, rax
    
    mov rsi, r12
    push rdi
    
    push 1
    pop rax
    mov edi, r8d
    
    mov dx, 2603
    syscall 
    
    push 3
    pop rax
    mov edi, r8d
    syscall
    
    pop rdi

    jmp short loop

jump_to_call:
    call server
    db "HTTP/1.0 200 OK", 13, 10
    db "Content-Type: text/html", 13, 10
    db "Content-Encoding: br", 13, 10, 13, 10
    
    incbin "index.html.br"

