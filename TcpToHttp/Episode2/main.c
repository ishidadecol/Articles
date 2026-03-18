// TcpToHttp/Episode2/main.c
#include <stdio.h>

int main(int argc, char *argv[]) {

  char buffer[9];
  size_t bytesRead;
  char filepath[] = "../message.txt";
  FILE *file = fopen(filepath, "r");

  if (file == NULL) {
    printf("Erro ao abrir o arquivo\n");
    return 1;
  }

  while ((bytesRead = fread(buffer, 1, 8, file)) > 0) {
    buffer[bytesRead] = '\0';
    printf("READ: %s\n", buffer);
  }

  fclose(file);
  return 0;
}
