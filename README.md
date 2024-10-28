# FCTube - Plataforma de Vídeos

**FCTube** é uma plataforma de vídeos inspirada no YouTube, desenvolvida com uma arquitetura de microserviços, oferecendo um painel administrativo robusto, processamento de mídia em nuvem e uma interface de usuário moderna. O projeto utiliza **Django** para o painel administrativo, **Golang** para o microserviço de processamento e **Next.js** para o front-end, com integração de mensagens via **RabbitMQ**.

## Visão Geral do Projeto

O projeto é dividido em três componentes principais:

1. **Painel Administrativo**: Permite o gerenciamento de vídeos, canais, usuários e estatísticas da plataforma.
2. **Microserviço de Compilação de Chunks**: Processa e compila vídeos em chunks, realiza o upload para a nuvem e gerencia tarefas de conversão de mídia.
3. **Front-end em Next.js**: Fornece uma interface para os usuários interagirem com a plataforma, visualizarem vídeos e explorarem canais.

## Arquitetura do Sistema

- **Comunicação entre Serviços**: Utilizamos o **RabbitMQ** para a comunicação entre o painel administrativo e o microserviço de compilação de chunks, garantindo integração em tempo real.
- **Processamento e Armazenamento em Nuvem**: O microserviço em Golang compila os chunks de vídeo e realiza o upload para a nuvem, tornando-os disponíveis para o front-end.
- **Front-end em Next.js**: A interface de usuário é responsiva e construída em **React** com **Next.js**, oferecendo navegação rápida e uma experiência de visualização otimizada.

## Funcionalidades

### Painel Administrativo
- **Gerenciamento de vídeos**: Upload, edição e exclusão de vídeos.
- **Controle de canais e usuários**: Administração de perfis e permissões.
- **Visualização de estatísticas**: Monitoramento de visualizações e interações.
- **Moderação de comentários** e conteúdo.
- Interface intuitiva e responsiva, desenvolvida em Django.

### Microserviço de Compilação de Chunks
- **Conversão de vídeos**: Processamento e compressão de vídeos em chunks.
- **Armazenamento em nuvem**: Upload de vídeos processados diretamente para o provedor de nuvem.
- **Mensageria com RabbitMQ**: Sincronização e envio de notificações ao painel administrativo sobre o status dos vídeos.

### Front-end (Next.js)
- **Interface de usuário**: Exploração de vídeos, canais e interações em tempo real.
- **Renderização SSR**: Renderização no lado do servidor para desempenho otimizado.
- **Reprodutor de vídeos**: Experiência de visualização com carregamento rápido.

## Tecnologias Utilizadas

- **Painel Administrativo**: Python, Django, PostgreSQL, RabbitMQ.
- **Microserviço de Compilação**: Golang, RabbitMQ.
- **Front-end**: JavaScript, React, Next.js.
- **Infraestrutura**: Docker para contêineres e configuração de ambientes.

## Instalação e Configuração

1. **Clone o repositório**:
   ```bash
   git clone https://github.com/seu-usuario/FCTube.git
