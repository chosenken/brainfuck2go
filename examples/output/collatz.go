package main
import "fmt"
func main() {
	buffer := make([]byte, 30000)
	ptr := 0
	ptr = ptr + 2
	buffer[ptr] = buffer[ptr] + byte(1)
	ptr = ptr + 1
	buffer[ptr] = buffer[ptr] + byte(1)
	ptr = ptr + -1
	for buffer[ptr] != 0 {
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
			ptr = ptr + 2
			for buffer[ptr] != 0 {
				ptr = ptr + 2
			}
			ptr = ptr + 3
			for buffer[ptr] != 0 {
				ptr = ptr + 2
			}
			buffer[ptr] = buffer[ptr] + byte(1)
			for buffer[ptr] != 0 {
				ptr = ptr + -2
			}
			ptr = ptr + -3
			for buffer[ptr] != 0 {
				ptr = ptr + -2
			}
			ptr = ptr + 1
			for buffer[ptr] != 0 {
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					ptr = ptr + 2
				}
				ptr = ptr + 2
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					ptr = ptr + 2
				}
				ptr = ptr + -1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + -1
				for buffer[ptr] != 0 {
					ptr = ptr + -2
				}
				ptr = ptr + -3
				for buffer[ptr] != 0 {
					ptr = ptr + -2
				}
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + 1
			for buffer[ptr] != 0 {
				ptr = ptr + 2
			}
			ptr = ptr + 2
			for buffer[ptr] != 0 {
				ptr = ptr + -4
				for buffer[ptr] != 0 {
					ptr = ptr + -2
				}
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					ptr = ptr + 2
				}
				ptr = ptr + 2
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + -4
			for buffer[ptr] != 0 {
				ptr = ptr + -2
			}
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr + 2
		}
		ptr = ptr + -2
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] + byte(5)
			for buffer[ptr] != 0 {
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] + byte(8)
				ptr = ptr + -1
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + 1
			fmt.Print(string(buffer[ptr]))
			ptr = ptr + -1
			buffer[ptr] = buffer[ptr] + byte(6)
			for buffer[ptr] != 0 {
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] - byte(8)
				ptr = ptr + -1
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr + -2
		}
		ptr = ptr + 2
		for buffer[ptr] != 0 {
			ptr = ptr + 2
		}
		buffer[ptr] = buffer[ptr] + byte(1)
		for buffer[ptr] != 0 {
			ptr = ptr + 4
			for buffer[ptr] != 0 {
				ptr = ptr + -2
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + -1
			buffer[ptr] = buffer[ptr] - byte(1)
			for buffer[ptr] != 0 {
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + -1
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr + -1
			for buffer[ptr] != 0 {
				ptr = ptr + -2
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 2
				buffer[ptr] = buffer[ptr] - byte(1)
				for buffer[ptr] != 0 {
					ptr = ptr + -2
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
				}
			}
			ptr = ptr + 3
			for buffer[ptr] != 0 {
				ptr = ptr + -3
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + -2
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 5
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + -3
			for buffer[ptr] != 0 {
				ptr = ptr + 3
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + -3
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + -2
			for buffer[ptr] != 0 {
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 2
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					ptr = ptr + -1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + -1
					for buffer[ptr] != 0 {
						ptr = ptr + -2
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 2
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + -1
					for buffer[ptr] != 0 {
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + -1
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + -1
					for buffer[ptr] != 0 {
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + -1
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 4
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + -1
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + -1
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + -1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						ptr = ptr + 2
					}
					ptr = ptr + -2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							ptr = ptr + -1
							buffer[ptr] = buffer[ptr] + byte(3)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] - byte(1)
							for buffer[ptr] != 0 {
								ptr = ptr + -1
								buffer[ptr] = buffer[ptr] + byte(3)
								ptr = ptr + 1
								buffer[ptr] = buffer[ptr] - byte(1)
								for buffer[ptr] != 0 {
									ptr = ptr + -1
									buffer[ptr] = buffer[ptr] + byte(3)
									ptr = ptr + 1
									buffer[ptr] = buffer[ptr] - byte(1)
									for buffer[ptr] != 0 {
										ptr = ptr + -1
										for buffer[ptr] != 0 {
											buffer[ptr] = buffer[ptr] - byte(1)
										}
										buffer[ptr] = buffer[ptr] + byte(2)
										ptr = ptr + 2
										for buffer[ptr] != 0 {
											buffer[ptr] = buffer[ptr] - byte(1)
										}
										buffer[ptr] = buffer[ptr] + byte(1)
										ptr = ptr + 1
										buffer[ptr] = buffer[ptr] + byte(1)
										ptr = ptr + -2
										buffer[ptr] = buffer[ptr] - byte(1)
										for buffer[ptr] != 0 {
											ptr = ptr + -1
											buffer[ptr] = buffer[ptr] + byte(3)
											ptr = ptr + 1
											buffer[ptr] = buffer[ptr] - byte(1)
											for buffer[ptr] != 0 {
												ptr = ptr + -1
												buffer[ptr] = buffer[ptr] + byte(3)
												ptr = ptr + 1
												buffer[ptr] = buffer[ptr] - byte(1)
												for buffer[ptr] != 0 {
													ptr = ptr + -1
													for buffer[ptr] != 0 {
														buffer[ptr] = buffer[ptr] - byte(1)
													}
													buffer[ptr] = buffer[ptr] + byte(1)
													ptr = ptr + 3
													buffer[ptr] = buffer[ptr] + byte(1)
													ptr = ptr + -2
													buffer[ptr] = buffer[ptr] - byte(1)
													for buffer[ptr] != 0 {
														ptr = ptr + -1
														buffer[ptr] = buffer[ptr] + byte(3)
														ptr = ptr + 1
														buffer[ptr] = buffer[ptr] - byte(1)
														for buffer[ptr] != 0 {
															ptr = ptr + -1
															buffer[ptr] = buffer[ptr] + byte(3)
															ptr = ptr + 1
															buffer[ptr] = buffer[ptr] - byte(1)
														}
													}
												}
											}
										}
									}
								}
							}
						}
						ptr = ptr + -1
						for buffer[ptr] != 0 {
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + -1
							buffer[ptr] = buffer[ptr] - byte(1)
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + -2
					}
					ptr = ptr + 3
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + -1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							ptr = ptr + -1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] - byte(1)
							for buffer[ptr] != 0 {
								ptr = ptr + -1
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 1
								buffer[ptr] = buffer[ptr] - byte(1)
								for buffer[ptr] != 0 {
									ptr = ptr + -1
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 1
									buffer[ptr] = buffer[ptr] - byte(1)
									for buffer[ptr] != 0 {
										ptr = ptr + -1
										buffer[ptr] = buffer[ptr] + byte(1)
										ptr = ptr + 1
										buffer[ptr] = buffer[ptr] - byte(1)
										for buffer[ptr] != 0 {
											ptr = ptr + -1
											buffer[ptr] = buffer[ptr] + byte(1)
											ptr = ptr + 1
											buffer[ptr] = buffer[ptr] - byte(1)
											for buffer[ptr] != 0 {
												ptr = ptr + -1
												buffer[ptr] = buffer[ptr] + byte(1)
												ptr = ptr + 1
												buffer[ptr] = buffer[ptr] - byte(1)
												for buffer[ptr] != 0 {
													ptr = ptr + -1
													buffer[ptr] = buffer[ptr] + byte(1)
													ptr = ptr + 1
													buffer[ptr] = buffer[ptr] - byte(1)
													for buffer[ptr] != 0 {
														ptr = ptr + -1
														buffer[ptr] = buffer[ptr] + byte(1)
														ptr = ptr + 1
														buffer[ptr] = buffer[ptr] - byte(1)
														for buffer[ptr] != 0 {
															ptr = ptr + -1
															buffer[ptr] = buffer[ptr] + byte(1)
															ptr = ptr + 1
															buffer[ptr] = buffer[ptr] - byte(1)
															for buffer[ptr] != 0 {
																ptr = ptr + -1
																for buffer[ptr] != 0 {
																	buffer[ptr] = buffer[ptr] - byte(1)
																}
																ptr = ptr + 2
																for buffer[ptr] != 0 {
																	buffer[ptr] = buffer[ptr] - byte(1)
																}
																buffer[ptr] = buffer[ptr] + byte(1)
																ptr = ptr + 1
																buffer[ptr] = buffer[ptr] + byte(1)
																ptr = ptr + -2
																buffer[ptr] = buffer[ptr] - byte(1)
																for buffer[ptr] != 0 {
																	ptr = ptr + -1
																	buffer[ptr] = buffer[ptr] + byte(1)
																	ptr = ptr + 1
																	buffer[ptr] = buffer[ptr] - byte(1)
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
						ptr = ptr + -1
						for buffer[ptr] != 0 {
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + -1
							buffer[ptr] = buffer[ptr] - byte(1)
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 2
					}
					ptr = ptr + -2
					for buffer[ptr] != 0 {
						ptr = ptr + -2
					}
					ptr = ptr + 1
				}
				ptr = ptr + -1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 1
					}
					ptr = ptr + -1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						for buffer[ptr] != 0 {
							ptr = ptr + -1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] - byte(1)
							for buffer[ptr] != 0 {
								ptr = ptr + -1
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + -1
								buffer[ptr] = buffer[ptr] - byte(1)
								for buffer[ptr] != 0 {
									ptr = ptr + -1
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 1
									buffer[ptr] = buffer[ptr] - byte(1)
									for buffer[ptr] != 0 {
										ptr = ptr + -1
										buffer[ptr] = buffer[ptr] - byte(1)
										ptr = ptr + 2
										buffer[ptr] = buffer[ptr] + byte(1)
										ptr = ptr + -1
										buffer[ptr] = buffer[ptr] - byte(1)
										for buffer[ptr] != 0 {
											ptr = ptr + -1
											buffer[ptr] = buffer[ptr] + byte(1)
											ptr = ptr + 1
											buffer[ptr] = buffer[ptr] - byte(1)
											for buffer[ptr] != 0 {
												ptr = ptr + -1
												buffer[ptr] = buffer[ptr] - byte(1)
												ptr = ptr + 2
												buffer[ptr] = buffer[ptr] + byte(1)
												ptr = ptr + -1
												buffer[ptr] = buffer[ptr] - byte(1)
												for buffer[ptr] != 0 {
													ptr = ptr + -1
													buffer[ptr] = buffer[ptr] + byte(1)
													ptr = ptr + 1
													buffer[ptr] = buffer[ptr] - byte(1)
													for buffer[ptr] != 0 {
														ptr = ptr + -1
														buffer[ptr] = buffer[ptr] - byte(1)
														ptr = ptr + 2
														buffer[ptr] = buffer[ptr] + byte(1)
														ptr = ptr + -1
														buffer[ptr] = buffer[ptr] - byte(1)
														for buffer[ptr] != 0 {
															ptr = ptr + -1
															buffer[ptr] = buffer[ptr] + byte(1)
															ptr = ptr + 1
															buffer[ptr] = buffer[ptr] - byte(1)
															for buffer[ptr] != 0 {
																ptr = ptr + -1
																buffer[ptr] = buffer[ptr] - byte(1)
																ptr = ptr + 2
																buffer[ptr] = buffer[ptr] + byte(1)
																ptr = ptr + -1
																buffer[ptr] = buffer[ptr] - byte(1)
																for buffer[ptr] != 0 {
																	ptr = ptr + -1
																	buffer[ptr] = buffer[ptr] + byte(1)
																	ptr = ptr + 1
																	buffer[ptr] = buffer[ptr] - byte(1)
																	for buffer[ptr] != 0 {
																		ptr = ptr + -1
																		buffer[ptr] = buffer[ptr] - byte(1)
																		ptr = ptr + 2
																		buffer[ptr] = buffer[ptr] + byte(1)
																		ptr = ptr + -1
																		buffer[ptr] = buffer[ptr] - byte(1)
																		for buffer[ptr] != 0 {
																			ptr = ptr + -1
																			buffer[ptr] = buffer[ptr] + byte(1)
																			ptr = ptr + 1
																			buffer[ptr] = buffer[ptr] - byte(1)
																			for buffer[ptr] != 0 {
																				ptr = ptr + -1
																				buffer[ptr] = buffer[ptr] - byte(1)
																				ptr = ptr + 2
																				buffer[ptr] = buffer[ptr] + byte(1)
																				ptr = ptr + -1
																				buffer[ptr] = buffer[ptr] - byte(1)
																				for buffer[ptr] != 0 {
																					ptr = ptr + -1
																					buffer[ptr] = buffer[ptr] + byte(1)
																					ptr = ptr + 1
																					buffer[ptr] = buffer[ptr] - byte(1)
																					for buffer[ptr] != 0 {
																						ptr = ptr + -1
																						buffer[ptr] = buffer[ptr] - byte(1)
																						ptr = ptr + 2
																						buffer[ptr] = buffer[ptr] + byte(1)
																						ptr = ptr + -1
																						buffer[ptr] = buffer[ptr] - byte(1)
																						for buffer[ptr] != 0 {
																							ptr = ptr + -1
																							buffer[ptr] = buffer[ptr] + byte(1)
																							ptr = ptr + 1
																							buffer[ptr] = buffer[ptr] - byte(1)
																							for buffer[ptr] != 0 {
																								ptr = ptr + -1
																								buffer[ptr] = buffer[ptr] - byte(1)
																								ptr = ptr + 2
																								buffer[ptr] = buffer[ptr] + byte(1)
																								ptr = ptr + -1
																								buffer[ptr] = buffer[ptr] - byte(1)
																								for buffer[ptr] != 0 {
																									ptr = ptr + -1
																									buffer[ptr] = buffer[ptr] + byte(1)
																									ptr = ptr + 1
																									buffer[ptr] = buffer[ptr] - byte(1)
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							ptr = ptr + -1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] - byte(1)
						}
						ptr = ptr + -1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + -1
						for buffer[ptr] != 0 {
							ptr = ptr + -1
							buffer[ptr] = buffer[ptr] + byte(10)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] - byte(1)
						}
						ptr = ptr + -1
					}
					ptr = ptr + 2
					for buffer[ptr] != 0 {
						ptr = ptr + -1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 2
					}
					ptr = ptr + -2
					for buffer[ptr] != 0 {
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + -2
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						ptr = ptr + -1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						ptr = ptr + -1
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
						}
					}
					ptr = ptr + -1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + -2
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + -2
					for buffer[ptr] != 0 {
						ptr = ptr + -2
					}
				}
				buffer[ptr] = buffer[ptr] + byte(6)
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] + byte(7)
					ptr = ptr + -1
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] + byte(2)
				fmt.Print(string(buffer[ptr]))
				buffer[ptr] = buffer[ptr] - byte(12)
				fmt.Print(string(buffer[ptr]))
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					ptr = ptr + 2
				}
				ptr = ptr + -2
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] + byte(5)
					for buffer[ptr] != 0 {
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] + byte(8)
						ptr = ptr + -1
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					fmt.Print(string(buffer[ptr]))
					ptr = ptr + -1
					buffer[ptr] = buffer[ptr] + byte(6)
					for buffer[ptr] != 0 {
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] - byte(8)
						ptr = ptr + -1
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + -2
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + -1
			}
			ptr = ptr + 1
			for buffer[ptr] != 0 {
				ptr = ptr + -1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + -1
		}
		ptr = ptr + 3
		for buffer[ptr] != 0 {
			ptr = ptr + 2
		}
		ptr = ptr + -2
		for buffer[ptr] != 0 {
			ptr = ptr + 1
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + -1
			buffer[ptr] = buffer[ptr] - byte(1)
			ptr = ptr + -2
		}
		buffer[ptr] = buffer[ptr] + byte(10)
		fmt.Print(string(buffer[ptr]))
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
		}
		ptr = ptr + -3
		for buffer[ptr] != 0 {
			ptr = ptr + -2
		}
		ptr = ptr + 3
		buffer[ptr] = buffer[ptr] + byte(1)
		ptr = ptr + -1
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
			ptr = ptr + 1
			for buffer[ptr] != 0 {
				ptr = ptr + -1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] - byte(1)
				for buffer[ptr] != 0 {
					ptr = ptr + -1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] - byte(1)
					for buffer[ptr] != 0 {
						ptr = ptr + -1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] - byte(1)
						for buffer[ptr] != 0 {
							ptr = ptr + -1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] - byte(1)
							for buffer[ptr] != 0 {
								ptr = ptr + -1
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 1
								buffer[ptr] = buffer[ptr] - byte(1)
								for buffer[ptr] != 0 {
									ptr = ptr + -1
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 1
									buffer[ptr] = buffer[ptr] - byte(1)
									for buffer[ptr] != 0 {
										ptr = ptr + -1
										buffer[ptr] = buffer[ptr] + byte(1)
										ptr = ptr + 1
										buffer[ptr] = buffer[ptr] - byte(1)
										for buffer[ptr] != 0 {
											ptr = ptr + -1
											buffer[ptr] = buffer[ptr] + byte(1)
											ptr = ptr + 1
											buffer[ptr] = buffer[ptr] - byte(1)
											for buffer[ptr] != 0 {
												ptr = ptr + -1
												buffer[ptr] = buffer[ptr] + byte(1)
												ptr = ptr + 1
												buffer[ptr] = buffer[ptr] - byte(1)
												for buffer[ptr] != 0 {
													ptr = ptr + -1
													for buffer[ptr] != 0 {
														buffer[ptr] = buffer[ptr] - byte(1)
													}
													ptr = ptr + 2
													for buffer[ptr] != 0 {
														buffer[ptr] = buffer[ptr] - byte(1)
													}
													buffer[ptr] = buffer[ptr] + byte(1)
													ptr = ptr + 1
													buffer[ptr] = buffer[ptr] + byte(1)
													ptr = ptr + -2
													buffer[ptr] = buffer[ptr] - byte(1)
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
			ptr = ptr + -1
			for buffer[ptr] != 0 {
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + -1
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr + 2
		}
		ptr = ptr + -2
		for buffer[ptr] != 0 {
			ptr = ptr + -2
		}
		ptr = ptr + 2
	}
}
