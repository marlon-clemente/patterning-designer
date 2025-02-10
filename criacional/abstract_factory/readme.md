# Abstract Factory

## 🏭 Factory Method vs Abstract Factory

| Padrão           | Objetivo                                                                       |
| ---------------- | ------------------------------------------------------------------------------ |
| Factory Method   | Cria um único tipo de produto, delegando a criação para subclasses.            |
| Abstract Factory | Cria famílias de produtos relacionados, garantindo compatibilidade entre eles. |

## 🎯 Factory Method → Foca na criação de um único objeto

📌 Quando usar?

Quando queremos delegar a criação de um objeto para subclasses, permitindo extensibilidade.

Exemplo (Factory Method)

Aqui, cada fábrica cria apenas um tipo de botão:

```go
    package main

    import "fmt"

    // Produto
    type Button interface {
        Render() string
    }

    // Produto Concreto 1
    type WindowsButton struct{}

    func (w WindowsButton) Render() string {
        return "Botão estilo Windows"
    }

    // Produto Concreto 2
    type MacButton struct{}

    func (m MacButton) Render() string {
        return "Botão estilo macOS"
    }

    // Factory Method
    type ButtonFactory interface {
        CreateButton() Button
    }

    // Concrete Factory 1
    type WindowsButtonFactory struct{}

    func (WindowsButtonFactory) CreateButton() Button {
        return WindowsButton{}
    }

    // Concrete Factory 2
    type MacButtonFactory struct{}

    func (MacButtonFactory) CreateButton() Button {
        return MacButton{}
    }

    // Cliente
    func main() {
        var factory ButtonFactory
        osType := "Mac" // Ou "Windows"

        if osType == "Windows" {
            factory = WindowsButtonFactory{}
        } else {
            factory = MacButtonFactory{}
        }

        button := factory.CreateButton()
        fmt.Println(button.Render()) // Output: Botão estilo macOS
    }

```

## 🏭🏭 Abstract Factory → Foca na criação de famílias de objetos

📌 Quando usar?

Quando precisamos criar múltiplos produtos relacionados que devem ser compatíveis entre si. (example_1.go).

| Característica | Factory Method                                   | Abstract Factory                                    |
| -------------- | ------------------------------------------------ | --------------------------------------------------- |
| Criação        | Um único produto.                                | Uma família de produtos relacionados                |
| Abstração      | Método abstrato para criar um objeto específico. | Interface para criar múltiplos produtos compatíveis |
| Uso comum      | Delegação da criação de um único objeto.         | Garantir compatibilidade entre objetos              |
| Exemplo        | Criar apenas botões.                             | Criar botões e checkboxes compatíveis               |
