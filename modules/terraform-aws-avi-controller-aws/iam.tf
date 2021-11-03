resource "aws_iam_role" "avi" {
  count              = var.create_iam ? 1 : 0
  name               = "${var.name_prefix}_AviController-Refined-Role"
  assume_role_policy = file("${path.module}/files/avicontroller-role-trust.json")

  tags = var.custom_tags
}
resource "aws_iam_instance_profile" "avi" {
  count = var.create_iam ? 1 : 0
  name  = "${var.name_prefix}_avi_iam_profile"
  role  = aws_iam_role.avi[0].id
}
resource "aws_iam_role_policy" "avi_ec2" {
  count = var.create_iam ? 1 : 0
  name  = "${var.name_prefix}-avicontroller-ec2-policy"
  role  = aws_iam_role.avi[0].id

  policy = file("${path.module}/files/avicontroller-ec2-policy.json")
}
resource "aws_iam_role_policy" "avi_iam" {
  count = var.create_iam ? 1 : 0
  name  = "${var.name_prefix}-avicontroller-iam-policy"
  role  = aws_iam_role.avi[0].id

  policy = file("${path.module}/files/avicontroller-iam-policy.json")
}
resource "aws_iam_role_policy" "avi_r53" {
  count = var.create_iam ? 1 : 0
  name  = "${var.name_prefix}-avicontroller-r53-policy"
  role  = aws_iam_role.avi[0].id

  policy = file("${path.module}/files/avicontroller-r53-policy.json")
}
resource "aws_iam_role_policy" "avi_autoscaling" {
  count = var.create_iam ? 1 : 0
  name  = "${var.name_prefix}-avicontroller-asg-policy"
  role  = aws_iam_role.avi[0].id

  policy = file("${path.module}/files/avicontroller-role-auto-scaling-group-policy.json")
}
resource "aws_iam_role_policy" "avi_sns" {
  count = var.create_iam ? 1 : 0
  name  = "${var.name_prefix}-avicontroller-sns-policy"
  role  = aws_iam_role.avi[0].id

  policy = file("${path.module}/files/avicontroller-sns-policy.json")
}
resource "aws_iam_role_policy" "avi_sqs" {
  count = var.create_iam ? 1 : 0
  name  = "${var.name_prefix}-avicontroller-sqs-policy"
  role  = aws_iam_role.avi[0].id

  policy = file("${path.module}/files/avicontroller-sqs-policy.json")
}
resource "aws_iam_role_policy" "avi_asg_notification" {
  count = var.create_iam ? 1 : 0
  name  = "${var.name_prefix}-avicontroller-asg-notification-policy"
  role  = aws_iam_role.avi[0].id

  policy = file("${path.module}/files/avicontroller-asg-notification-policy.json")
}
resource "aws_iam_role_policy" "avi_kms" {
  count = var.create_iam ? 1 : 0
  name  = "${var.name_prefix}-avicontroller-kms-policy"
  role  = aws_iam_role.avi[0].id

  policy = file("${path.module}/files/avicontroller-asg-notification-policy.json")
}