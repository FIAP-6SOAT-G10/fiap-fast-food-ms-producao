Feature: Payment Succeed

  Scenario: Mensagem publicada no broker de pagamento
    Given uma mensagem de pagamento chega no broker SQS
    When o pagamento é aprovado
    Then uma mensagem de confirmação é publicada no sistema
