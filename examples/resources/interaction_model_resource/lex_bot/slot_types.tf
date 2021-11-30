resource "aws_lex_slot_type" "qna_slot_type" {

  name                     = "QnaSlotType"
  value_selection_strategy = "ORIGINAL_VALUE"
  
  description = "All custom questions that have canned responses"

  enumeration_value {
    value = "how do I reset my password?"
  }

  enumeration_value {
    value = "I forgot my password"
  }

  enumeration_value {
    value = "Can't remember my password"
  }

  enumeration_value {
    value = "My login does not work"
  }

  enumeration_value {
    value = "help my gas is leaking"
  }

  enumeration_value {
    value = "emergency leak"
  }

  enumeration_value {
    value = "I smell gas in my house"
  }

  enumeration_value {
    value = "what should i do if my pilot light is out?"
  }

    enumeration_value {
    value = "pilot light"
  }

    enumeration_value {
    value = "help with pilot light on furnace"
  }

    enumeration_value {
    value = "exit"
  }

    enumeration_value {
    value = "quit"
  }  
}
