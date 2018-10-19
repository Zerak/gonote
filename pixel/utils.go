package main

func CalculateWithdrawalAndUpdate(cnv *Canvas, player *Account, grd *Guard) int {
	patronBonus := cnv.PatronBonusScaled(player) / PRECISION_BASE

	balance := player.balanceScaled / PRECISION_BASE

	withdrawAmount := patronBonus + balance
	// eosio_assert(withdrawAmount > 0, "Balance too small for withdrawal");

	// eosio_assert(grd.quota >= withdrawAmount, "Contract withdrawal quota exceeded");

	grd.quota -= withdrawAmount

	// Find the actual truncated values, and use those to update the records
	withdrawnBonusScaled := patronBonus * PRECISION_BASE
	withdrawnBalanceScaled := balance * PRECISION_BASE

	// Due to precision issues in PRECISION_BASE, withdrawnBalanceScaled  may cause overflow.
	if withdrawnBalanceScaled >= player.balanceScaled {
		player.balanceScaled = 0
	} else {
		player.balanceScaled -= withdrawnBalanceScaled
	}
	player.maskScaled += withdrawnBonusScaled

	return withdrawAmount
}
