package awslex


// Provides a prompt for making sure that the user is ready for the intent to be fulfilled.
//
// Example:
//
//
type CfnBot_IntentConfirmationSettingProperty struct {
	// Prompts the user to confirm the intent. This question should have a yes or no answer.
	//
	// Amazon Lex uses this prompt to ensure that the user acknowledges that the intent is ready for fulfillment. For example, with the `OrderPizza` intent, you might want to confirm that the order is correct before placing it. For other intents, such as intents that simply respond to user questions, you might not need to ask the user for confirmation before providing the information.
	PromptSpecification interface{} `field:"required" json:"promptSpecification" yaml:"promptSpecification"`
	// The `DialogCodeHookInvocationSetting` object associated with intent's confirmation step.
	//
	// The dialog code hook is triggered based on these invocation settings when the confirmation next step or declination next step or failure next step is `InvokeDialogCodeHook` .
	CodeHook interface{} `field:"optional" json:"codeHook" yaml:"codeHook"`
	// A list of conditional branches to evaluate after the intent is closed.
	ConfirmationConditional interface{} `field:"optional" json:"confirmationConditional" yaml:"confirmationConditional"`
	// Specifies the next step that the bot executes when the customer confirms the intent.
	ConfirmationNextStep interface{} `field:"optional" json:"confirmationNextStep" yaml:"confirmationNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond the user input.
	ConfirmationResponse interface{} `field:"optional" json:"confirmationResponse" yaml:"confirmationResponse"`
	// A list of conditional branches to evaluate after the intent is declined.
	DeclinationConditional interface{} `field:"optional" json:"declinationConditional" yaml:"declinationConditional"`
	// Specifies the next step that the bot executes when the customer declines the intent.
	DeclinationNextStep interface{} `field:"optional" json:"declinationNextStep" yaml:"declinationNextStep"`
	// When the user answers "no" to the question defined in `promptSpecification` , Amazon Lex responds with this response to acknowledge that the intent was canceled.
	DeclinationResponse interface{} `field:"optional" json:"declinationResponse" yaml:"declinationResponse"`
	// The `DialogCodeHookInvocationSetting` used when the code hook is invoked during confirmation prompt retries.
	ElicitationCodeHook interface{} `field:"optional" json:"elicitationCodeHook" yaml:"elicitationCodeHook"`
	// Provides a list of conditional branches.
	//
	// Branches are evaluated in the order that they are entered in the list. The first branch with a condition that evaluates to true is executed. The last branch in the list is the default branch. The default branch should not have any condition expression. The default branch is executed if no other branch has a matching condition.
	FailureConditional interface{} `field:"optional" json:"failureConditional" yaml:"failureConditional"`
	// The next step to take in the conversation if the confirmation step fails.
	FailureNextStep interface{} `field:"optional" json:"failureNextStep" yaml:"failureNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond the user input when the intent confirmation fails.
	FailureResponse interface{} `field:"optional" json:"failureResponse" yaml:"failureResponse"`
	// Specifies whether the intent's confirmation is sent to the user.
	//
	// When this field is false, confirmation and declination responses aren't sent. If the `IsActive` field isn't specified, the default is true.
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
}

