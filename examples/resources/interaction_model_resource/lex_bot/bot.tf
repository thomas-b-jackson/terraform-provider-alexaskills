resource "aws_lex_bot" "socal_qna" {

  locale           = "en-US"
  name             = "SoCalQnA"
  process_behavior = "SAVE"
  voice_id         = "Salli"
    
  abort_statement {
    message {
      content_type = "PlainText"
      content      = "Sorry, I am not able to assist at this time"
    }
  }

  child_directed = false

  clarification_prompt {
    max_attempts = 2

    message {
      content_type = "PlainText"
      content      = "Sorry, what can I help you with?"
    }
  }

  description                 = "Bot to answer simple customer q-n-a questions"
  detect_sentiment            = false
  idle_session_ttl_in_seconds = 600

  intent {
    intent_name    = aws_lex_intent.qna_intent.name
    intent_version = aws_lex_intent.qna_intent.version
  }
}
