package awslex


// Defines the action that the bot executes at runtime when the conversation reaches this step.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dialogActionProperty := &DialogActionProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	SlotToElicit: jsii.String("slotToElicit"),
//   	SuppressNextMessage: jsii.Boolean(false),
//   }
//
type CfnBot_DialogActionProperty struct {
	// The action that the bot should execute.
	//
	// Valid values are `ElicitIntent` , `StartIntent` , `ElicitSlot` , `EvaluateConditional` , `InvokeDialogCodeHook` , `ConfirmIntent` , `FulfillIntent` , `CloseIntent` , and `EndConversation` .
	Type *string `field:"required" json:"type" yaml:"type"`
	// If the dialog action is `ElicitSlot` , defines the slot to elicit from the user.
	SlotToElicit *string `field:"optional" json:"slotToElicit" yaml:"slotToElicit"`
	// When true the next message for the intent is not used.
	SuppressNextMessage interface{} `field:"optional" json:"suppressNextMessage" yaml:"suppressNextMessage"`
}
