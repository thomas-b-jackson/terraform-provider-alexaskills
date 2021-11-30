output "bot_questions" {
  value = resource.aws_lex_slot_type.qna_slot_type.enumeration_value
}

output "bot_name" {
  value = resource.aws_lex_slot_type.qna_slot_type.name
}