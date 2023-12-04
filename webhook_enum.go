// Code generated by paypalc. DO NOT EDIT.
package paypal

// EventType is the type of webhook.
//
// See https://developer.paypal.com/api/rest/webhooks/event-names/
type EventType string


// Payments
//
// The webhooks for authorizing and capturing payments correspond to both supported versions of the
// Payments API

// V1

const (
	// A payment authorization is created, approved, executed, or a future payment authorization is
	// created.
	//
	// Related method: [Create payment] with `intent` set to `authorize`
	//
	// [Create payment]: https://developer.paypal.com/docs/api/payments/v1/#payment_create
	// (redeclared) PaymentAuthorizationCreated EventType = "PAYMENT.AUTHORIZATION.CREATED"

	// A payment authorization is voided.
	//
	// Related method: [Void authorization]
	//
	// [Void authorization]: https://developer.paypal.com/docs/api/payments/v1/#authorization_void
	// (redeclared) PaymentAuthorizationVoided EventType = "PAYMENT.AUTHORIZATION.VOIDED"

	// A payment capture completes.
	//
	// Related method: [Show captured payment details]
	//
	// [Show captured payment details]: https://developer.paypal.com/docs/api/payments/v1/#capture_get
	// (redeclared) PaymentCaptureCompleted EventType = "PAYMENT.CAPTURE.COMPLETED"

	// A payment capture is denied.
	//
	// Related method: [Show captured payment details]
	//
	// [Show captured payment details]: https://developer.paypal.com/docs/api/payments/v1/#capture_get
	PaymentCaptureDenied EventType = "PAYMENT.CAPTURE.DENIED"

	// The state of a payment capture changes to pending.
	//
	// Related method: [Show captured payment details]
	//
	// [Show captured payment details]: https://developer.paypal.com/docs/api/payments/v1/#capture_get
	// (redeclared) PaymentCapturePending EventType = "PAYMENT.CAPTURE.PENDING"

	// A merchant refunds a payment capture.
	//
	// Related method: [Refund captured payment]
	//
	// [Refund captured payment]: https://developer.paypal.com/docs/api/payments/v1/#capture_refund
	// (redeclared) PaymentCaptureRefunded EventType = "PAYMENT.CAPTURE.REFUNDED"

	// PayPal reverses a payment capture.
	//
	// Related method: [Refund captured payment]
	//
	// [Refund captured payment]: https://developer.paypal.com/docs/api/payments/v1/#capture_refund
	// (redeclared) PaymentCaptureReversed EventType = "PAYMENT.CAPTURE.REVERSED"
)

// V2

const (
	// A payment authorization is created, approved, executed, or a future payment authorization is
	// created.
	//
	// Related method: [Capture authorized payment]
	//
	// [Capture authorized payment]: https://developer.paypal.com/docs/api/payments/v2/#authorizations_capture
	PaymentAuthorizationCreated EventType = "PAYMENT.AUTHORIZATION.CREATED"

	// A payment authorization is voided either due to authorization reaching it’s 30 day validity
	// period or authorization was manually voided using the Void Authorized Payment API.
	//
	// Related method: [Show details for authorized payment] with response `status` of `voided`.
	//
	// [Show details for authorized payment]: https://developer.paypal.com/docs/api/payments/v2/#authorizations_get
	PaymentAuthorizationVoided EventType = "PAYMENT.AUTHORIZATION.VOIDED"

	// A payment capture is declined.
	//
	// Related method: [Capture authorized payment] with `status` of `declined`.
	//
	// [Capture authorized payment]: https://developer.paypal.comhttps://developer.paypal.com/docs/api/payments/v2/#authorizations_capture
	PaymentCaptureDeclined EventType = "PAYMENT.CAPTURE.DECLINED"

	// A payment capture completes.
	//
	// Related method: [Capture authorized payment] with response `status` of `completed`.
	//
	// [Capture authorized payment]: https://developer.paypal.com/docs/api/payments/v2/#authorizations_capture
	PaymentCaptureCompleted EventType = "PAYMENT.CAPTURE.COMPLETED"

	// The state of a payment capture changes to pending.
	//
	// Related method: [Show details for authorized payment] with response `status` of `pending`.
	//
	// [Show details for authorized payment]: https://developer.paypal.com/docs/api/payments/v2/#authorizations_capture
	PaymentCapturePending EventType = "PAYMENT.CAPTURE.PENDING"

	// A merchant refunds a payment capture.
	//
	// Related method: [Capture authorized payment] with response `status` of `refunded`.
	//
	// [Capture authorized payment]: https://developer.paypal.com/docs/api/payments/v2/#authorizations_capture
	PaymentCaptureRefunded EventType = "PAYMENT.CAPTURE.REFUNDED"

	// PayPal reverses a payment capture.
	//
	// Related method: [Refund captured payment]
	//
	// [Refund captured payment]: https://developer.paypal.com/docs/api/payments/v2/#captures_refund
	PaymentCaptureReversed EventType = "PAYMENT.CAPTURE.REVERSED"
)


// Batch payouts
//
// To get item-related information, use the [HATEOAS links] from the webhooks response.
//
// [HATEOAS links]: https://developer.paypal.com/api/rest/responses/#hateoas-links

const (
	// A batch payout payment is denied.
	//
	// Related method: [Show payout details]
	//
	// [Show payout details]: https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts_get
	PaymentPayoutsbatchDenied EventType = "PAYMENT.PAYOUTSBATCH.DENIED"

	// The state of a batch payout payment changes to processing.
	//
	// Related method: [Show payout details]
	//
	// [Show payout details]: https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts_get
	PaymentPayoutsbatchProcessing EventType = "PAYMENT.PAYOUTSBATCH.PROCESSING"

	// A batch payout payment completes successfully.
	//
	// Related method: [Show payout details]
	//
	// [Show payout details]: https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts_get
	PaymentPayoutsbatchSuccess EventType = "PAYMENT.PAYOUTSBATCH.SUCCESS"

	// A payouts item is blocked.
	//
	// Related method: [Show payout item details]
	//
	// [Show payout item details]: https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-item_get
	PaymentPayoutsItemBlocked EventType = "PAYMENT.PAYOUTS-ITEM.BLOCKED"

	// A payouts item is canceled.
	//
	// Related method: [Cancel unclaimed payout item]
	//
	// [Cancel unclaimed payout item]: https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-item_cancel
	PaymentPayoutsItemCanceled EventType = "PAYMENT.PAYOUTS-ITEM.CANCELED"

	// A payouts item is denied.
	//
	// Related method: [Show payout item details]
	//
	// [Show payout item details]: https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-item_get
	PaymentPayoutsItemDenied EventType = "PAYMENT.PAYOUTS-ITEM.DENIED"

	// A payouts item fails.
	//
	// Related method: [Show payout item details]
	//
	// [Show payout item details]: https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-item_get
	PaymentPayoutsItemFailed EventType = "PAYMENT.PAYOUTS-ITEM.FAILED"

	// A payouts item is held.
	//
	// Related method: [Show payout item details]
	//
	// [Show payout item details]: https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-item_get
	PaymentPayoutsItemHeld EventType = "PAYMENT.PAYOUTS-ITEM.HELD"

	// A payouts item is refunded.
	//
	// Related method: [Show payout item details]
	//
	// [Show payout item details]: https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-item_get
	PaymentPayoutsItemRefunded EventType = "PAYMENT.PAYOUTS-ITEM.REFUNDED"

	// A payouts item is returned.
	//
	// Related method: [Show payout item details]
	//
	// [Show payout item details]: https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-item_get
	PaymentPayoutsItemReturned EventType = "PAYMENT.PAYOUTS-ITEM.RETURNED"

	// A payouts item succeeds.
	//
	// Related method: [Show payout item details]
	//
	// [Show payout item details]: https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-item_get
	PaymentPayoutsItemSucceeded EventType = "PAYMENT.PAYOUTS-ITEM.SUCCEEDED"

	// A payouts item is unclaimed.
	//
	// Related method: [Show payout item details]
	//
	// [Show payout item details]: https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-item_get
	PaymentPayoutsItemUnclaimed EventType = "PAYMENT.PAYOUTS-ITEM.UNCLAIMED"
)


// Billing plans and agreements
//
// Deprecated: The billing plans and agreements webhooks are deprecated along with the [Billing
// Agreements REST API].
//
// [Billing Agreements REST API]: https://developer.paypal.com/docs/api/payments.billing-agreements/v1

const (
	// A billing plan is created.
	//
	// Related method: [Create billing plan]
	//
	// [Create billing plan]: https://developer.paypal.com/docs/api/payments.billing-plans/v1/#billing-plans_post
	BillingPlanCreated EventType = "BILLING.PLAN.CREATED"

	// A billing plan is updated.
	//
	// Related method: [Update billing plan]
	//
	// [Update billing plan]: https://developer.paypal.com/docs/api/payments.billing-plans/v1/#billing-plans_patch
	BillingPlanUpdated EventType = "BILLING.PLAN.UPDATED"

	// A billing agreement is canceled.
	//
	// Related method: [Cancel agreement]
	//
	// [Cancel agreement]: https://developer.paypal.com/docs/api/payments.billing-agreements/v1/#billing-agreements_cancel
	BillingSubscriptionCancelled EventType = "BILLING.SUBSCRIPTION.CANCELLED"

	// A billing agreement is created.
	//
	// Related method: [Create agreement]
	//
	// [Create agreement]: https://developer.paypal.com/docs/api/payments.billing-agreements/v1/#billing-agreements_re-activate
	BillingSubscriptionCreated EventType = "BILLING.SUBSCRIPTION.CREATED"

	// A billing agreement is re-activated.
	//
	// Related method: [Re-activate agreement]
	//
	// [Re-activate agreement]: https://developer.paypal.com/docs/api/payments.billing-agreements/v1/#billing-agreements_re-activate
	BillingSubscriptionReActivated EventType = "BILLING.SUBSCRIPTION.RE-ACTIVATED"

	// A billing agreement is suspended.
	//
	// Related method: [Suspend agreement]
	//
	// [Suspend agreement]: https://developer.paypal.com/docs/api/payments.billing-agreements/v1/#billing-agreements_suspend
	BillingSubscriptionSuspended EventType = "BILLING.SUBSCRIPTION.SUSPENDED"

	// A billing agreement is updated.
	//
	// Related method: [Update agreement]
	//
	// [Update agreement]: https://developer.paypal.com/docs/api/payments.billing-agreements/v1/#billing-agreements_patch
	BillingSubscriptionUpdated EventType = "BILLING.SUBSCRIPTION.UPDATED"
)


// Log in with PayPal

const (
	// A user's consent token is revoked.
	//
	// Related method: Not applicable
	IdentityAuthorizationConsentRevoked EventType = "IDENTITY.AUTHORIZATION-CONSENT.REVOKED"
)


// Checkout buyer approval

const (
	// Checkout payment is created and approved by buyer.
	//
	// Related method: [Payment details]
	//
	// [Payment details]: https://developer.paypal.com/docs/api/payments/v1/#payment_get
	PaymentsPaymentCreated EventType = "PAYMENTS.PAYMENT.CREATED"

	// A buyer approved a checkout order
	//
	// Related method: [Orders]
	//
	// [Orders]: https://developer.paypal.com/api/rest/webhooks/event-names/#orders
	CheckoutOrderApproved EventType = "CHECKOUT.ORDER.APPROVED"

	// Express checkout payment is created and approved by buyer.
	//
	// Related method: [Express Checkout]
	//
	// [Express Checkout]: https://developer.paypal.com/api/nvp-soap/payflow/express-checkout/sale/#the-express-checkout-basic-integration
	CheckoutCheckoutBuyerApproved EventType = "CHECKOUT.CHECKOUT.BUYER-APPROVED"
)


// Disputes

const (
	// A dispute is created.
	//
	// Related method: N/A
	CustomerDisputeCreated EventType = "CUSTOMER.DISPUTE.CREATED"

	// A dispute is resolved.
	//
	// Related method: [Settle dispute]
	//
	// [Settle dispute]: https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes-actions_adjudicate
	CustomerDisputeResolved EventType = "CUSTOMER.DISPUTE.RESOLVED"

	// A dispute is updated.
	//
	// Related method: [Partially update dispute]
	//
	// [Partially update dispute]: https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes_patch
	CustomerDisputeUpdated EventType = "CUSTOMER.DISPUTE.UPDATED"

	// A risk dispute is created.
	//
	// Related method: The `CUSTOMER.DISPUTE.CREATED` event type supersedes and deprecates this
	// event type.
	RiskDisputeCreated EventType = "RISK.DISPUTE.CREATED"
)


// Invoicing

const (
	// A merchant or customer cancels an invoice.
	//
	// Related method: [Cancel invoice]
	//
	// [Cancel invoice]: https://developer.paypal.com/docs/api/invoicing/v1/#invoices_get
	InvoicingInvoiceCancelled EventType = "INVOICING.INVOICE.CANCELLED"

	// An invoice is created.
	//
	// Related method: [Create draft invoice]
	//
	// [Create draft invoice]: https://developer.paypal.com/docs/api/invoicing/v1/#invoices_create
	InvoicingInvoiceCreated EventType = "INVOICING.INVOICE.CREATED"

	// An invoice is paid, partially paid, or payment is made and is pending.
	//
	// Related method: [Mark invoice as paid]
	//
	// [Mark invoice as paid]: https://developer.paypal.com/docs/api/invoicing/v1/#invoices_record-payment
	InvoicingInvoicePaid EventType = "INVOICING.INVOICE.PAID"

	// An invoice is refunded or partially refunded.
	//
	// Related method: [Mark invoice as refunded]
	//
	// [Mark invoice as refunded]: https://developer.paypal.com/docs/api/invoicing/v1/#invoices_record-refund
	InvoicingInvoiceRefunded EventType = "INVOICING.INVOICE.REFUNDED"

	// An invoice is scheduled.
	//
	// Related method: [Schedule invoice]
	//
	// [Schedule invoice]: https://developer.paypal.com/docs/api/invoicing/v1/#invoices_schedule
	InvoicingInvoiceScheduled EventType = "INVOICING.INVOICE.SCHEDULED"

	// An invoice is updated.
	//
	// Related method: [Update invoice]
	//
	// [Update invoice]: https://developer.paypal.com/docs/api/invoicing/v1/#invoices_update
	InvoicingInvoiceUpdated EventType = "INVOICING.INVOICE.UPDATED"
)


// Marketplaces and Platforms

const (
	// See [CheckoutOrderCompleted].
	// (redeclared) CheckoutOrderCompleted EventType = "CHECKOUT.ORDER.COMPLETED"

	// See [CheckoutOrderApproved].
	// (redeclared) CheckoutOrderApproved EventType = "CHECKOUT.ORDER.APPROVED"

	// See [CheckoutOrderProcessed].
	// (redeclared) CheckoutOrderProcessed EventType = "CHECKOUT.ORDER.PROCESSED"

	// A limitation is added for a partner's managed account.
	//
	// Related method: [Update managed account]
	//
	// [Update managed account]: https://developer.paypal.com/api/limited-release/managed-accounts/v3
	CustomerAccountLimitationAdded EventType = "CUSTOMER.ACCOUNT-LIMITATION.ADDED"

	// A limitation is escalated for a partner's managed account.
	//
	// Related method: [Update managed account]
	//
	// [Update managed account]: https://developer.paypal.com/api/limited-release/managed-accounts/v3
	CustomerAccountLimitationEscalated EventType = "CUSTOMER.ACCOUNT-LIMITATION.ESCALATED"

	// A limitation is lifted for a partner's managed account.
	//
	// Related method: [Update managed account]
	//
	// [Update managed account]: https://developer.paypal.com/api/limited-release/managed-accounts/v3
	CustomerAccountLimitationLifted EventType = "CUSTOMER.ACCOUNT-LIMITATION.LIFTED"

	// A limitation is updated for a partner's managed account.
	//
	// Related method: [Update managed account]
	//
	// [Update managed account]: https://developer.paypal.com/api/limited-release/managed-accounts/v3
	CustomerAccountLimitationUpdated EventType = "CUSTOMER.ACCOUNT-LIMITATION.UPDATED"

	// PayPal must enable the merchant's account as `PPCP` for this webhook to work.
	//
	// Related method: [Create partner referral]
	//
	// [Create partner referral]: https://developer.paypal.com/docs/api/partner-referrals/v2/#partner-referrals_create
	CustomerMerchantIntegrationCapabilityUpdated EventType = "CUSTOMER.MERCHANT-INTEGRATION.CAPABILITY-UPDATED"

	// The products available to the merchant have changed.
	//
	// Related method: [Create partner referral]
	//
	// [Create partner referral]: https://developer.paypal.com/docs/api/partner-referrals/v2/#partner-referrals_create
	CustomerMerchantIntegrationProductSubscriptionUpdated EventType = "CUSTOMER.MERCHANT-INTEGRATION.PRODUCT-SUBSCRIPTION-UPDATED"

	// Merchant onboards again to a partner.
	//
	// Related method: [Create partner referral]
	//
	// [Create partner referral]: https://developer.paypal.com/docs/api/partner-referrals/v2/#partner-referrals_create
	CustomerMerchantIntegrationSellerAlreadyIntegrated EventType = "CUSTOMER.MERCHANT-INTEGRATION.SELLER-ALREADY-INTEGRATED"

	// PayPal creates a merchant account from the partner's onboarding link.
	//
	// Related method: [Create partner referral]
	//
	// [Create partner referral]: https://developer.paypal.com/docs/api/partner-referrals/v2/#partner-referrals_create
	CustomerMerchantIntegrationSellerOnboardingInitiated EventType = "CUSTOMER.MERCHANT-INTEGRATION.SELLER-ONBOARDING-INITIATED"

	// Merchant grants consents to a partner.
	//
	// Related method: [Create partner referral]
	//
	// [Create partner referral]: https://developer.paypal.com/docs/api/partner-referrals/v2/#partner-referrals_create
	CustomerMerchantIntegrationSellerConsentGranted EventType = "CUSTOMER.MERCHANT-INTEGRATION.SELLER-CONSENT-GRANTED"

	// Merchant confirms the email and consents are granted.
	//
	// Related method: [Create partner referral]
	//
	// [Create partner referral]: https://developer.paypal.com/docs/api/partner-referrals/v2/#partner-referrals_create
	CustomerMerchantIntegrationSellerEmailConfirmed EventType = "CUSTOMER.MERCHANT-INTEGRATION.SELLER-EMAIL-CONFIRMED"

	// Merchant completes setup.
	//
	// Related method: See `MERCHANT.ONBOARDING.COMPLETED` in [Merchant onboarding]
	//
	// [Merchant onboarding]: https://developer.paypal.com#merchant-onboarding
	MerchantOnboardingCompleted EventType = "MERCHANT.ONBOARDING.COMPLETED"

	// The consents for a merchant account setup are revoked or an account is closed.
	//
	// Related method: See `MERCHANT.PARTNER-CONSENT.REVOKED` in [Merchant onboarding]
	//
	// [Merchant onboarding]: https://developer.paypal.com#merchant-onboarding
	MerchantPartnerConsentRevoked EventType = "MERCHANT.PARTNER-CONSENT.REVOKED"

	// A payment capture completes.
	//
	// Related method: See `PAYMENT.CAPTURE.COMPLETED` in [Payments]
	//
	// [Payments]: https://developer.paypal.com#authorized-and-captured-payments
	// (redeclared) PaymentCaptureCompleted EventType = "PAYMENT.CAPTURE.COMPLETED"

	// A payment capture is denied.
	//
	// Related method: See `PAYMENT.CAPTURE.DENIED` in [Payments]
	//
	// [Payments]: https://developer.paypal.com#authorized-and-captured-payments
	// (redeclared) PaymentCaptureDenied EventType = "PAYMENT.CAPTURE.DENIED"

	// A merchant refunds a payment capture.
	//
	// Related method: See `PAYMENT.CAPTURE.REFUNDED` in [Payments]
	//
	// [Payments]: https://developer.paypal.com#authorized-and-captured-payments
	// (redeclared) PaymentCaptureRefunded EventType = "PAYMENT.CAPTURE.REFUNDED"

	// Funds are disbursed to the seller and partner.
	//
	// Related method: See `PAYMENT.REFERENCED-PAYOUT-ITEM.COMPLETE` in [Payments]
	//
	// [Payments]: https://developer.paypal.com#authorized-and-captured-payments
	PaymentReferencedPayoutItemCompleted EventType = "PAYMENT.REFERENCED-PAYOUT-ITEM.COMPLETED"

	// Attempt to disburse funds fails.
	//
	// Related method: See `PAYMENT.REFERENCED-PAYOUT-ITEM.FAILED` in [Payments]
	//
	// [Payments]: https://developer.paypal.com#authorized-and-captured-payments
	PaymentReferencedPayoutItemFailed EventType = "PAYMENT.REFERENCED-PAYOUT-ITEM.FAILED"
)


// Merchant onboarding

const (
	// A merchant completes setup.
	//
	// Related method: [Seller Onboarding]
	//
	// [Seller Onboarding]: https://developer.paypal.com/docs/multiparty/seller-onboarding/
	// (redeclared) MerchantOnboardingCompleted EventType = "MERCHANT.ONBOARDING.COMPLETED"

	// The consents for a merchant account setup are revoked or an account is closed.
	//
	// Related method: [Seller Onboarding]
	//
	// [Seller Onboarding]: https://developer.paypal.com/docs/multiparty/seller-onboarding/
	// (redeclared) MerchantPartnerConsentRevoked EventType = "MERCHANT.PARTNER-CONSENT.REVOKED"

	// Managed account has been created.
	//
	// Related method: [Managed Path Onboarding]
	//
	// [Managed Path Onboarding]: https://developer.paypal.com/limited-release/commerce-platform/v3/seller-onboarding/managed-seller-onboarding
	CustomerManagedAccountAccountCreated EventType = "CUSTOMER.MANAGED-ACCOUNT.ACCOUNT-CREATED"

	// Managed account creation failed.
	//
	// Related method: [Managed Path Onboarding]
	//
	// [Managed Path Onboarding]: https://developer.paypal.com/limited-release/commerce-platform/v3/seller-onboarding/managed-seller-onboarding
	CustomerManagedAccountCreationFailed EventType = "CUSTOMER.MANAGED-ACCOUNT.CREATION-FAILED"

	// Managed account has been updated.
	//
	// Related method: [Managed Path Onboarding]
	//
	// [Managed Path Onboarding]: https://developer.paypal.com/limited-release/commerce-platform/v3/seller-onboarding/managed-seller-onboarding
	CustomerManagedAccountAccountUpdated EventType = "CUSTOMER.MANAGED-ACCOUNT.ACCOUNT-UPDATED"

	// Capabilities and/or process status has been changed on a managed account.
	//
	// Related method: [Managed Path Onboarding]
	//
	// [Managed Path Onboarding]: https://developer.paypal.com/limited-release/commerce-platform/v3/seller-onboarding/managed-seller-onboarding
	CustomerManagedAccountAccountStatusChanged EventType = "CUSTOMER.MANAGED-ACCOUNT.ACCOUNT-STATUS-CHANGED"

	// Managed account has been risk assessed or the risk assessment has been changed.
	//
	// Related method: [Managed Path Onboarding]
	//
	// [Managed Path Onboarding]: https://developer.paypal.com/limited-release/commerce-platform/v3/seller-onboarding/managed-seller-onboarding
	CustomerManagedAccountRiskAssessed EventType = "CUSTOMER.MANAGED-ACCOUNT.RISK-ASSESSED"

	// Negative balance debit has been notified on a managed account.
	//
	// Related method: [Update managed account]
	//
	// [Update managed account]: https://developer.paypal.com/api/limited-release/managed-accounts/v3
	CustomerManagedAccountNegativeBalanceNotified EventType = "CUSTOMER.MANAGED-ACCOUNT.NEGATIVE-BALANCE-NOTIFIED"

	// Negative balance debit has been initiated on a managed account.
	//
	// Related method: [Update managed account]
	//
	// [Update managed account]: https://developer.paypal.com/api/limited-release/managed-accounts/v3
	CustomerManagedAccountNegativeBalanceDebitInitiated EventType = "CUSTOMER.MANAGED-ACCOUNT.NEGATIVE-BALANCE-DEBIT-INITIATED"

	// A limitation was added for given merchant account.
	//
	// Related method: [Update managed account]
	//
	// [Update managed account]: https://developer.paypal.com/api/limited-release/managed-accounts/v3
	// (redeclared) CustomerAccountLimitationAdded EventType = "CUSTOMER.ACCOUNT-LIMITATION.ADDED"

	// A limitation was lifted for given merchant account.
	//
	// Related method: [Update managed account]
	//
	// [Update managed account]: https://developer.paypal.com/api/limited-release/managed-accounts/v3
	// (redeclared) CustomerAccountLimitationLifted EventType = "CUSTOMER.ACCOUNT-LIMITATION.LIFTED"

	// A limitation was updated for given merchant account.
	//
	// Related method: [Update managed account]
	//
	// [Update managed account]: https://developer.paypal.com/api/limited-release/managed-accounts/v3
	// (redeclared) CustomerAccountLimitationUpdated EventType = "CUSTOMER.ACCOUNT-LIMITATION.UPDATED"

	// A limitation was escalated for given merchant account.
	//
	// Related method: [Update managed account]
	//
	// [Update managed account]: https://developer.paypal.com/api/limited-release/managed-accounts/v3
	// (redeclared) CustomerAccountLimitationEscalated EventType = "CUSTOMER.ACCOUNT-LIMITATION.ESCALATED"
)


// Orders
//
// The webhooks for orders correspond to both supported versions of the Orders API

// V1

const (
	// A checkout order is processed.
	//
	// Related method: [Orders API]
	//
	// [Orders API]: https://developer.paypal.com/docs/api/orders/v1/
	CheckoutOrderProcessed EventType = "CHECKOUT.ORDER.PROCESSED"
)

// V2

const (
	// A checkout order is processed. <strong>Note:</strong> For use by marketplaces and platforms
	// only.
	//
	// Related method: [Orders API]
	//
	// [Orders API]: https://developer.paypal.com/docs/api/orders/v2/
	CheckoutOrderCompleted EventType = "CHECKOUT.ORDER.COMPLETED"

	// A buyer approved a checkout order
	//
	// Related method: [Orders API]
	//
	// [Orders API]: https://developer.paypal.com/docs/api/orders/v2/
	// (redeclared) CheckoutOrderApproved EventType = "CHECKOUT.ORDER.APPROVED"

	// A problem occurred after the buyer approved the order but before you captured the payment.
	// Refer to [Handle uncaptured payments] for what to do when this event occurs.
	//
	// Related method: [Orders API]
	//
	// [Handle uncaptured payments]: https://developer.paypal.com/docs/checkout/apm/reference/handle-uncaptured-payments/
	// [Orders API]: https://developer.paypal.com/docs/api/orders/v2/
	CheckoutPaymentApprovalReversed EventType = "CHECKOUT.PAYMENT-APPROVAL.REVERSED"
)


// Payment orders

const (
	// A payment order is canceled.
	//
	// Related method: [Void order]
	//
	// [Void order]: https://developer.paypal.com/docs/api/payments/v1/#orders_void
	PaymentOrderCancelled EventType = "PAYMENT.ORDER.CANCELLED"

	// A payment order is created.
	//
	// Related method: [Create payment]
	//
	// [Create payment]: https://developer.paypal.com/docs/api/payments/v2/#authorizations_capture
	PaymentOrderCreated EventType = "PAYMENT.ORDER.CREATED"
)


// Referenced payouts

const (
	// Funds are disbursed to the seller and partner.
	//
	// Related method: [Create referenced payout item]
	//
	// [Create referenced payout item]: https://developer.paypal.com/docs/api/referenced-payouts/v1/#referenced-payouts-items_create
	// (redeclared) PaymentReferencedPayoutItemCompleted EventType = "PAYMENT.REFERENCED-PAYOUT-ITEM.COMPLETED"

	// Attempt to disburse funds fails.
	//
	// Related method: [Create referenced payout item]
	//
	// [Create referenced payout item]: https://developer.paypal.com/docs/api/referenced-payouts/v1/#referenced-payouts-items_create
	// (redeclared) PaymentReferencedPayoutItemFailed EventType = "PAYMENT.REFERENCED-PAYOUT-ITEM.FAILED"
)


// Sales

const (
	// A sale completes.
	//
	// Related method: [Show sale details]
	//
	// [Show sale details]: https://developer.paypal.com/docs/api/payments/v1/#sale_get
	PaymentSaleCompleted EventType = "PAYMENT.SALE.COMPLETED"

	// The state of a sale changes from pending to denied.
	//
	// Related method: [Show sale details]
	//
	// [Show sale details]: https://developer.paypal.com/docs/api/payments/v1/#sale_get
	PaymentSaleDenied EventType = "PAYMENT.SALE.DENIED"

	// The state of a sale changes to pending.
	//
	// Related method: [Show sale details]
	//
	// [Show sale details]: https://developer.paypal.com/docs/api/payments/v1/#sale_get
	PaymentSalePending EventType = "PAYMENT.SALE.PENDING"

	// A merchant refunds a sale.
	//
	// Related method: [Refund sale]
	//
	// [Refund sale]: https://developer.paypal.com/docs/api/payments/v1/#sale_refund
	PaymentSaleRefunded EventType = "PAYMENT.SALE.REFUNDED"

	// PayPal reverses a sale.
	//
	// Related method: [Refund sale]
	//
	// [Refund sale]: https://developer.paypal.com/docs/api/payments/v1/#sale_refund
	PaymentSaleReversed EventType = "PAYMENT.SALE.REVERSED"
)


// Subscriptions

const (
	// A product is created.
	//
	// Related method: [Create product]
	//
	// [Create product]: https://developer.paypal.com/docs/api/catalog-products/v1/#products_create
	CatalogProductCreated EventType = "CATALOG.PRODUCT.CREATED"

	// A product is updated.
	//
	// Related method: [Update product]
	//
	// [Update product]: https://developer.paypal.com/docs/api/catalog-products/v1/#products_patch
	CatalogProductUpdated EventType = "CATALOG.PRODUCT.UPDATED"

	// A payment is made on a subscription.
	// (redeclared) PaymentSaleCompleted EventType = "PAYMENT.SALE.COMPLETED"

	// A merchant refunds a sale.
	// (redeclared) PaymentSaleRefunded EventType = "PAYMENT.SALE.REFUNDED"

	// A payment is reversed on a subscription.
	// (redeclared) PaymentSaleReversed EventType = "PAYMENT.SALE.REVERSED"

	// A billing plan is created.
	//
	// Related method: [Create plan]
	//
	// [Create plan]: https://developer.paypal.com/docs/api/subscriptions/v1/#plans_create
	// (redeclared) BillingPlanCreated EventType = "BILLING.PLAN.CREATED"

	// A billing plan is updated.
	//
	// Related method: [Update plan]
	//
	// [Update plan]: https://developer.paypal.com/docs/api/subscriptions/v1/#plans_patch
	// (redeclared) BillingPlanUpdated EventType = "BILLING.PLAN.UPDATED"

	// A billing plan is activated.
	//
	// Related method: [Activate plan]
	//
	// [Activate plan]: https://developer.paypal.com/docs/api/subscriptions/v1/#plans_activate
	BillingPlanActivated EventType = "BILLING.PLAN.ACTIVATED"

	// A price change for the plan is activated.
	//
	// Related method: [Update pricing]
	//
	// [Update pricing]: https://developer.paypal.com/docs/api/subscriptions/v1/#plans_update-pricing-schemes
	BillingPlanPricingChangeActivated EventType = "BILLING.PLAN.PRICING-CHANGE.ACTIVATED"

	// A billing plan is deactivated.
	//
	// Related method: [Deactivate plan]
	//
	// [Deactivate plan]: https://developer.paypal.com/docs/api/subscriptions/v1/#plans_deactivate
	BillingPlanDeactivated EventType = "BILLING.PLAN.DEACTIVATED"

	// A subscription is created.
	//
	// Related method: [Create subscription]
	//
	// [Create subscription]: https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_create
	// (redeclared) BillingSubscriptionCreated EventType = "BILLING.SUBSCRIPTION.CREATED"

	// A subscription is activated.
	//
	// Related method: [Activate subscription]
	//
	// [Activate subscription]: https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_activate
	BillingSubscriptionActivated EventType = "BILLING.SUBSCRIPTION.ACTIVATED"

	// A subscription is updated.
	//
	// Related method: [Update subscription]
	//
	// [Update subscription]: https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_patch
	// (redeclared) BillingSubscriptionUpdated EventType = "BILLING.SUBSCRIPTION.UPDATED"

	// A subscription expires.
	//
	// Related method: [Show subscription details]
	//
	// [Show subscription details]: https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_get
	BillingSubscriptionExpired EventType = "BILLING.SUBSCRIPTION.EXPIRED"

	// A subscription is cancelled.
	//
	// Related method: [Cancel subscription]
	//
	// [Cancel subscription]: https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_cancel
	// (redeclared) BillingSubscriptionCancelled EventType = "BILLING.SUBSCRIPTION.CANCELLED"

	// A subscription is suspended.
	//
	// Related method: [Suspend subscription]
	//
	// [Suspend subscription]: https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_suspend
	// (redeclared) BillingSubscriptionSuspended EventType = "BILLING.SUBSCRIPTION.SUSPENDED"

	// Payment failed on subscription.
	//
	// Related method: [Subscription failed payment details]
	//
	// [Subscription failed payment details]: https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions-get-response
	BillingSubscriptionPaymentFailed EventType = "BILLING.SUBSCRIPTION.PAYMENT.FAILED"
)


// Vault

const (
	// A credit card is created.
	//
	// Related method: Store credit card
	VaultCreditCardCreated EventType = "VAULT.CREDIT-CARD.CREATED"

	// A credit card is deleted.
	//
	// Related method: Delete vaulted credit card
	VaultCreditCardDeleted EventType = "VAULT.CREDIT-CARD.DELETED"

	// A credit card is updated.
	//
	// Related method: Update vaulted credit card
	VaultCreditCardUpdated EventType = "VAULT.CREDIT-CARD.UPDATED"
)


// Additional information

