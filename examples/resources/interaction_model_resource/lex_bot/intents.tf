resource "aws_lex_intent" "qna_intent" {
  
  depends_on = [resource.aws_lex_slot_type.qna_slot_type]
  
  name = "QnaIntent"

  confirmation_prompt {
    max_attempts = 2

    message {
      content      = "Good to go"
      content_type = "PlainText"
    }
  }

  description = "qna"

  sample_utterances = [
    "{qnaslot}"
  ]

  fulfillment_activity {
    type = "ReturnIntent"
  }

  rejection_statement {
    message {
      content      = "Okay, never mind"
      content_type = "PlainText"
    }
  }

  slot {
    description = "Answer a question included in the slot"
    name        = "qnaslot"
    priority    = 1

    slot_constraint = "Required"
    slot_type       = "QnaSlotType"
    slot_type_version = "$LATEST"

    value_elicitation_prompt {
      max_attempts = 2

      message {
        content      = "What is the question?"
        content_type = "PlainText"
      }
    }
  }
}