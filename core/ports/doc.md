- As portas (ports) contém as assinaturas dos métodos que são utilizados pelos adaptadores, a fim de realizar as operações desejadas.

- Os casos de uso (use cases) representam a implementação da lógica de negócio, independente do tipo de banco de dados utilizado, ou de como o serviço será exposto (http ou grpc, por exemplo). Nesse momento, podemos observar a utilização das portas secundárias, que são responsáveis por exportar os repositórios.
