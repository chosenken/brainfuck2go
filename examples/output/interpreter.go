package main
import "fmt"
import "os"
func main() {
	buffer := make([]byte, 30000000)
	ptr := 0
	b := make([]byte, 1)
	ptr = ptr + 4
	buffer[ptr] = buffer[ptr] + byte(1)
	for buffer[ptr] != 0 {
		buffer[ptr] = buffer[ptr] - byte(1)
		ptr = ptr + 3
		buffer[ptr] = buffer[ptr] + byte(2)
		ptr = ptr + 1
		buffer[ptr] = buffer[ptr] + byte(1)
		ptr = ptr + 1
		buffer[ptr] = buffer[ptr] + byte(7)
		for buffer[ptr] != 0 {
			ptr = ptr - 1
			buffer[ptr] = buffer[ptr] + byte(4)
			ptr = ptr + 2
			buffer[ptr] = buffer[ptr] + byte(2)
			ptr = ptr - 1
			buffer[ptr] = buffer[ptr] - byte(1)
		}
		buffer[ptr] = buffer[ptr] + byte(2)
		ptr = ptr + 2
		buffer[ptr] = buffer[ptr] + byte(1)
		ptr = ptr + 1
		buffer[ptr] = buffer[ptr] + byte(1)
		ptr = ptr + 1
		buffer[ptr] = buffer[ptr] + byte(5)
		for buffer[ptr] != 0 {
			ptr = ptr + 1
			buffer[ptr] = buffer[ptr] + byte(2)
			ptr = ptr + 1
			buffer[ptr] = buffer[ptr] + byte(6)
			ptr = ptr - 2
			buffer[ptr] = buffer[ptr] - byte(1)
		}
		buffer[ptr] = buffer[ptr] + byte(1)
		ptr = ptr + 3
		os.Stdin.Read(b)
		buffer[ptr] = b[0]
		ptr = ptr - 1
		buffer[ptr] = buffer[ptr] + byte(2)
		for buffer[ptr] != 0 {
			for buffer[ptr] != 0 {
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 2
				}
				ptr = ptr - 1
				for buffer[ptr] != 0 {
					ptr = ptr + 2
				}
				ptr = ptr - 2
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr - 1
			for buffer[ptr] != 0 {
				ptr = ptr - 1
			}
			ptr = ptr - 1
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr + 2
			for buffer[ptr] != 0 {
				ptr = ptr + 1
			}
			ptr = ptr + 1
			for buffer[ptr] != 0 {
				ptr = ptr - 1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] - byte(1)
				for buffer[ptr] != 0 {
					for buffer[ptr] != 0 {
						ptr = ptr - 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
				}
				ptr = ptr - 1
				for buffer[ptr] != 0 {
					for buffer[ptr] != 0 {
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
						}
						ptr = ptr - 1
					}
					buffer[ptr] = buffer[ptr] + byte(2)
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] - byte(1)
					for buffer[ptr] != 0 {
						ptr = ptr - 1
						buffer[ptr] = buffer[ptr] + byte(9)
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							ptr = ptr - 1
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] - byte(1)
						}
						ptr = ptr + 2
					}
					ptr = ptr + 2
				}
			}
			ptr = ptr - 2
		}
		ptr = ptr + 1
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
		}
		buffer[ptr] = buffer[ptr] + byte(1)
		ptr = ptr - 2
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(2)
			for buffer[ptr] != 0 {
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 2
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr - 3
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 1
			}
			ptr = ptr + 2
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr - 4
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 2
				buffer[ptr] = buffer[ptr] + byte(1)
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						ptr = ptr - 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 6
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 5
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr - 1
					for buffer[ptr] != 0 {
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 7
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									for buffer[ptr] != 0 {
										buffer[ptr] = buffer[ptr] - byte(1)
										for buffer[ptr] != 0 {
											buffer[ptr] = buffer[ptr] - byte(1)
											for buffer[ptr] != 0 {
												buffer[ptr] = buffer[ptr] - byte(1)
												for buffer[ptr] != 0 {
													buffer[ptr] = buffer[ptr] - byte(1)
													for buffer[ptr] != 0 {
														for buffer[ptr] != 0 {
															buffer[ptr] = buffer[ptr] - byte(1)
														}
														ptr = ptr + 1
														buffer[ptr] = buffer[ptr] - byte(1)
														ptr = ptr - 1
													}
													ptr = ptr + 1
													for buffer[ptr] != 0 {
														buffer[ptr] = buffer[ptr] - byte(1)
														ptr = ptr - 2
														buffer[ptr] = buffer[ptr] + byte(7)
														ptr = ptr + 2
													}
													ptr = ptr - 1
												}
											}
										}
										ptr = ptr + 1
										for buffer[ptr] != 0 {
											buffer[ptr] = buffer[ptr] - byte(1)
										}
										ptr = ptr - 1
									}
									ptr = ptr + 1
									for buffer[ptr] != 0 {
										buffer[ptr] = buffer[ptr] - byte(1)
										ptr = ptr - 2
										buffer[ptr] = buffer[ptr] + byte(3)
										ptr = ptr + 2
									}
									ptr = ptr - 1
								}
								ptr = ptr + 1
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr - 2
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 2
								}
								ptr = ptr - 1
							}
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
							}
							ptr = ptr - 1
						}
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 7
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 7
						}
						ptr = ptr - 1
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr - 5
				}
				ptr = ptr + 3
				for buffer[ptr] != 0 {
					ptr = ptr - 2
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 2
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr - 1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									for buffer[ptr] != 0 {
										buffer[ptr] = buffer[ptr] - byte(1)
										for buffer[ptr] != 0 {
											buffer[ptr] = buffer[ptr] - byte(1)
											for buffer[ptr] != 0 {
												buffer[ptr] = buffer[ptr] - byte(1)
												for buffer[ptr] != 0 {
													buffer[ptr] = buffer[ptr] - byte(1)
													for buffer[ptr] != 0 {
														buffer[ptr] = buffer[ptr] - byte(1)
														for buffer[ptr] != 0 {
															for buffer[ptr] != 0 {
																buffer[ptr] = buffer[ptr] - byte(1)
															}
															ptr = ptr + 1
															buffer[ptr] = buffer[ptr] - byte(1)
															ptr = ptr - 3
															for buffer[ptr] != 0 {
																buffer[ptr] = buffer[ptr] - byte(1)
															}
															ptr = ptr - 2
															buffer[ptr] = buffer[ptr] + byte(1)
															ptr = ptr + 2
														}
													}
												}
											}
											ptr = ptr + 1
											for buffer[ptr] != 0 {
												buffer[ptr] = buffer[ptr] - byte(1)
											}
											ptr = ptr - 1
										}
										ptr = ptr + 1
										for buffer[ptr] != 0 {
											buffer[ptr] = buffer[ptr] - byte(1)
											ptr = ptr - 3
											for buffer[ptr] != 0 {
												buffer[ptr] = buffer[ptr] - byte(1)
											}
											ptr = ptr - 2
											buffer[ptr] = buffer[ptr] + byte(7)
											ptr = ptr + 3
										}
										ptr = ptr - 1
									}
								}
							}
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
							}
							ptr = ptr - 1
						}
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 3
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
							}
							ptr = ptr - 2
							buffer[ptr] = buffer[ptr] + byte(7)
							ptr = ptr + 3
						}
						ptr = ptr - 1
					}
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr - 6
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 2
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 3
					for buffer[ptr] != 0 {
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 2
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 3
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						ptr = ptr - 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							for buffer[ptr] != 0 {
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
								}
								ptr = ptr + 2
							}
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								ptr = ptr - 2
								for buffer[ptr] != 0 {
									ptr = ptr - 1
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 2
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr - 1
									buffer[ptr] = buffer[ptr] - byte(1)
								}
								ptr = ptr + 1
								for buffer[ptr] != 0 {
									ptr = ptr - 1
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 1
									buffer[ptr] = buffer[ptr] - byte(1)
								}
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 2
								buffer[ptr] = buffer[ptr] - byte(1)
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									for buffer[ptr] != 0 {
										buffer[ptr] = buffer[ptr] - byte(1)
										for buffer[ptr] != 0 {
											buffer[ptr] = buffer[ptr] - byte(1)
											for buffer[ptr] != 0 {
												buffer[ptr] = buffer[ptr] - byte(1)
												for buffer[ptr] != 0 {
													buffer[ptr] = buffer[ptr] - byte(1)
													for buffer[ptr] != 0 {
														buffer[ptr] = buffer[ptr] - byte(1)
														for buffer[ptr] != 0 {
															buffer[ptr] = buffer[ptr] - byte(1)
															for buffer[ptr] != 0 {
																buffer[ptr] = buffer[ptr] - byte(1)
																for buffer[ptr] != 0 {
																	buffer[ptr] = buffer[ptr] - byte(1)
																	for buffer[ptr] != 0 {
																		buffer[ptr] = buffer[ptr] - byte(1)
																		ptr = ptr - 1
																		buffer[ptr] = buffer[ptr] + byte(1)
																		ptr = ptr + 1
																	}
																	ptr = ptr - 1
																	buffer[ptr] = buffer[ptr] + byte(13)
																	ptr = ptr + 2
																	for buffer[ptr] != 0 {
																		buffer[ptr] = buffer[ptr] - byte(1)
																	}
																	ptr = ptr + 1
																	buffer[ptr] = buffer[ptr] - byte(1)
																	ptr = ptr + 1
																	buffer[ptr] = buffer[ptr] - byte(1)
																	ptr = ptr - 2
																}
															}
														}
														ptr = ptr + 2
														for buffer[ptr] != 0 {
															buffer[ptr] = buffer[ptr] - byte(1)
															ptr = ptr + 5
														}
														ptr = ptr - 2
													}
													ptr = ptr + 2
													for buffer[ptr] != 0 {
														buffer[ptr] = buffer[ptr] - byte(1)
														ptr = ptr - 3
														buffer[ptr] = buffer[ptr] + byte(12)
														ptr = ptr + 2
														for buffer[ptr] != 0 {
															buffer[ptr] = buffer[ptr] - byte(1)
														}
														ptr = ptr + 2
														buffer[ptr] = buffer[ptr] - byte(1)
													}
													ptr = ptr - 2
												}
												ptr = ptr + 2
												for buffer[ptr] != 0 {
													buffer[ptr] = buffer[ptr] - byte(1)
													ptr = ptr + 5
												}
												ptr = ptr - 2
											}
											ptr = ptr + 2
											for buffer[ptr] != 0 {
												buffer[ptr] = buffer[ptr] - byte(1)
												ptr = ptr - 3
												buffer[ptr] = buffer[ptr] + byte(11)
												ptr = ptr + 2
												for buffer[ptr] != 0 {
													buffer[ptr] = buffer[ptr] - byte(1)
												}
												ptr = ptr + 2
												buffer[ptr] = buffer[ptr] - byte(1)
											}
											ptr = ptr - 2
										}
										ptr = ptr + 2
										for buffer[ptr] != 0 {
											buffer[ptr] = buffer[ptr] - byte(1)
											ptr = ptr - 3
											buffer[ptr] = buffer[ptr] + byte(10)
											ptr = ptr + 2
											for buffer[ptr] != 0 {
												buffer[ptr] = buffer[ptr] - byte(1)
											}
											ptr = ptr + 2
											buffer[ptr] = buffer[ptr] - byte(1)
										}
										ptr = ptr - 2
									}
								}
								ptr = ptr + 2
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr + 5
								}
								ptr = ptr - 2
							}
							ptr = ptr - 1
						}
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						ptr = ptr + 2
					}
					ptr = ptr - 3
				}
				ptr = ptr + 2
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 3
			}
			ptr = ptr - 2
		}
		ptr = ptr + 2
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
			ptr = ptr + 1
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr + 1
		}
		ptr = ptr - 3
	}
	ptr = ptr - 2
	for buffer[ptr] != 0 {
		ptr = ptr - 2
	}
	ptr = ptr + 2
	for buffer[ptr] != 0 {
		for buffer[ptr] != 0 {
			ptr = ptr - 1
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr + 2
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr - 1
			buffer[ptr] = buffer[ptr] - byte(1)
		}
		buffer[ptr] = buffer[ptr] + byte(1)
		ptr = ptr - 1
		buffer[ptr] = buffer[ptr] - byte(1)
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									for buffer[ptr] != 0 {
										buffer[ptr] = buffer[ptr] - byte(1)
										for buffer[ptr] != 0 {
											buffer[ptr] = buffer[ptr] - byte(1)
											for buffer[ptr] != 0 {
												buffer[ptr] = buffer[ptr] - byte(1)
												for buffer[ptr] != 0 {
													buffer[ptr] = buffer[ptr] - byte(1)
													for buffer[ptr] != 0 {
														buffer[ptr] = buffer[ptr] - byte(1)
														for buffer[ptr] != 0 {
															buffer[ptr] = buffer[ptr] - byte(1)
															for buffer[ptr] != 0 {
																buffer[ptr] = buffer[ptr] - byte(1)
																for buffer[ptr] != 0 {
																	buffer[ptr] = buffer[ptr] - byte(1)
																	ptr = ptr + 1
																	buffer[ptr] = buffer[ptr] - byte(1)
																	ptr = ptr + 2
																	for buffer[ptr] != 0 {
																		ptr = ptr + 2
																	}
																	ptr = ptr + 2
																	for buffer[ptr] != 0 {
																		ptr = ptr + 2
																	}
																	ptr = ptr - 1
																	for buffer[ptr] != 0 {
																		ptr = ptr - 1
																		buffer[ptr] = buffer[ptr] - byte(1)
																		ptr = ptr - 2
																		buffer[ptr] = buffer[ptr] - byte(1)
																		ptr = ptr - 1
																	}
																	ptr = ptr - 1
																	for buffer[ptr] != 0 {
																		ptr = ptr - 2
																	}
																	ptr = ptr - 2
																	for buffer[ptr] != 0 {
																		ptr = ptr - 2
																	}
																	ptr = ptr - 1
																}
																ptr = ptr + 1
																for buffer[ptr] != 0 {
																	buffer[ptr] = buffer[ptr] - byte(1)
																	ptr = ptr + 2
																	for buffer[ptr] != 0 {
																		ptr = ptr + 2
																	}
																	ptr = ptr + 2
																	for buffer[ptr] != 0 {
																		ptr = ptr + 2
																	}
																	ptr = ptr - 1
																	for buffer[ptr] != 0 {
																		ptr = ptr + 1
																		buffer[ptr] = buffer[ptr] + byte(1)
																		ptr = ptr + 2
																		buffer[ptr] = buffer[ptr] + byte(1)
																		ptr = ptr + 1
																	}
																	ptr = ptr - 1
																	for buffer[ptr] != 0 {
																		ptr = ptr - 2
																	}
																	ptr = ptr - 2
																	for buffer[ptr] != 0 {
																		ptr = ptr - 2
																	}
																}
																ptr = ptr - 1
															}
															ptr = ptr + 1
															for buffer[ptr] != 0 {
																buffer[ptr] = buffer[ptr] - byte(1)
																ptr = ptr + 2
																for buffer[ptr] != 0 {
																	ptr = ptr + 2
																}
																ptr = ptr + 2
																for buffer[ptr] != 0 {
																	ptr = ptr + 2
																}
																ptr = ptr - 1
																for buffer[ptr] != 0 {
																	buffer[ptr] = buffer[ptr] - byte(1)
																}
																ptr = ptr - 1
																for buffer[ptr] != 0 {
																	ptr = ptr - 2
																}
																ptr = ptr - 2
																for buffer[ptr] != 0 {
																	ptr = ptr - 2
																}
															}
															ptr = ptr - 1
														}
														ptr = ptr + 1
														for buffer[ptr] != 0 {
															buffer[ptr] = buffer[ptr] - byte(1)
															ptr = ptr + 2
															for buffer[ptr] != 0 {
																ptr = ptr + 2
															}
															ptr = ptr + 2
															for buffer[ptr] != 0 {
																ptr = ptr + 2
															}
															ptr = ptr - 1
															for buffer[ptr] != 0 {
																ptr = ptr - 1
																buffer[ptr] = buffer[ptr] - byte(1)
																ptr = ptr - 1
															}
															ptr = ptr - 1
															for buffer[ptr] != 0 {
																ptr = ptr - 2
															}
															ptr = ptr - 2
															for buffer[ptr] != 0 {
																ptr = ptr - 2
															}
														}
														ptr = ptr - 1
													}
													ptr = ptr + 1
													for buffer[ptr] != 0 {
														buffer[ptr] = buffer[ptr] - byte(1)
														ptr = ptr + 2
														for buffer[ptr] != 0 {
															ptr = ptr + 2
														}
														ptr = ptr + 2
														for buffer[ptr] != 0 {
															ptr = ptr + 2
														}
														ptr = ptr - 1
														for buffer[ptr] != 0 {
															ptr = ptr + 1
															buffer[ptr] = buffer[ptr] + byte(1)
															ptr = ptr + 1
														}
														ptr = ptr - 1
														for buffer[ptr] != 0 {
															ptr = ptr - 2
														}
														ptr = ptr - 2
														for buffer[ptr] != 0 {
															ptr = ptr - 2
														}
													}
													ptr = ptr - 1
												}
												ptr = ptr + 1
												for buffer[ptr] != 0 {
													buffer[ptr] = buffer[ptr] - byte(1)
													ptr = ptr + 2
													for buffer[ptr] != 0 {
														ptr = ptr + 2
													}
													ptr = ptr + 2
													for buffer[ptr] != 0 {
														ptr = ptr + 2
													}
													ptr = ptr - 2
													buffer[ptr] = buffer[ptr] - byte(1)
													ptr = ptr - 2
													buffer[ptr] = buffer[ptr] - byte(1)
													ptr = ptr - 2
													for buffer[ptr] != 0 {
														ptr = ptr - 2
													}
													ptr = ptr - 2
													for buffer[ptr] != 0 {
														ptr = ptr - 2
													}
												}
												ptr = ptr - 1
											}
											ptr = ptr + 1
											for buffer[ptr] != 0 {
												buffer[ptr] = buffer[ptr] - byte(1)
												ptr = ptr + 2
												for buffer[ptr] != 0 {
													ptr = ptr + 2
												}
												ptr = ptr + 2
												for buffer[ptr] != 0 {
													ptr = ptr + 2
												}
												buffer[ptr] = buffer[ptr] + byte(1)
												ptr = ptr + 2
												buffer[ptr] = buffer[ptr] + byte(1)
												for buffer[ptr] != 0 {
													ptr = ptr - 2
												}
												ptr = ptr - 2
												for buffer[ptr] != 0 {
													ptr = ptr - 2
												}
											}
											ptr = ptr - 1
										}
										ptr = ptr + 1
										for buffer[ptr] != 0 {
											buffer[ptr] = buffer[ptr] - byte(1)
											ptr = ptr + 2
											for buffer[ptr] != 0 {
												ptr = ptr + 2
											}
											ptr = ptr + 2
											for buffer[ptr] != 0 {
												ptr = ptr + 2
											}
											ptr = ptr - 1
											buffer[ptr] = buffer[ptr] + byte(2)
											ptr = ptr - 1
											for buffer[ptr] != 0 {
												ptr = ptr - 2
											}
											ptr = ptr - 2
											for buffer[ptr] != 0 {
												ptr = ptr - 2
											}
										}
										ptr = ptr - 1
									}
									ptr = ptr + 1
									for buffer[ptr] != 0 {
										buffer[ptr] = buffer[ptr] - byte(1)
										ptr = ptr + 2
										for buffer[ptr] != 0 {
											ptr = ptr + 2
										}
										ptr = ptr + 2
										for buffer[ptr] != 0 {
											ptr = ptr + 2
										}
										ptr = ptr - 1
										buffer[ptr] = buffer[ptr] + byte(1)
										ptr = ptr - 1
										for buffer[ptr] != 0 {
											ptr = ptr - 2
										}
										ptr = ptr - 2
										for buffer[ptr] != 0 {
											ptr = ptr - 2
										}
									}
									ptr = ptr - 1
								}
								ptr = ptr + 1
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr + 2
									for buffer[ptr] != 0 {
										ptr = ptr + 2
									}
									ptr = ptr + 2
									for buffer[ptr] != 0 {
										ptr = ptr + 2
									}
									ptr = ptr - 1
									os.Stdin.Read(b)
									buffer[ptr] = b[0]
									ptr = ptr - 1
									for buffer[ptr] != 0 {
										ptr = ptr - 2
									}
									ptr = ptr - 2
									for buffer[ptr] != 0 {
										ptr = ptr - 2
									}
								}
								ptr = ptr - 1
							}
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 2
								for buffer[ptr] != 0 {
									ptr = ptr + 2
								}
								ptr = ptr + 2
								for buffer[ptr] != 0 {
									ptr = ptr + 2
								}
								ptr = ptr - 1
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 1
								for buffer[ptr] != 0 {
									ptr = ptr - 2
								}
								ptr = ptr - 2
								for buffer[ptr] != 0 {
									ptr = ptr - 2
								}
							}
							ptr = ptr - 1
						}
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 2
							for buffer[ptr] != 0 {
								ptr = ptr + 2
							}
							ptr = ptr + 2
							for buffer[ptr] != 0 {
								ptr = ptr + 2
							}
							ptr = ptr - 1
							fmt.Print(string(buffer[ptr]))
							ptr = ptr - 1
							for buffer[ptr] != 0 {
								ptr = ptr - 2
							}
							ptr = ptr - 2
							for buffer[ptr] != 0 {
								ptr = ptr - 2
							}
						}
						ptr = ptr - 1
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 2
						for buffer[ptr] != 0 {
							ptr = ptr + 2
						}
						ptr = ptr + 2
						for buffer[ptr] != 0 {
							ptr = ptr + 2
						}
						ptr = ptr - 2
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 2
						for buffer[ptr] != 0 {
							ptr = ptr - 2
						}
						ptr = ptr - 2
						for buffer[ptr] != 0 {
							ptr = ptr - 2
						}
					}
					ptr = ptr - 1
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 2
					for buffer[ptr] != 0 {
						ptr = ptr + 2
					}
					ptr = ptr + 2
					for buffer[ptr] != 0 {
						ptr = ptr + 2
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					for buffer[ptr] != 0 {
						ptr = ptr - 2
					}
					ptr = ptr - 2
					for buffer[ptr] != 0 {
						ptr = ptr - 2
					}
				}
				ptr = ptr - 1
			}
			ptr = ptr + 1
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 2
				for buffer[ptr] != 0 {
					ptr = ptr + 2
				}
				ptr = ptr + 2
				for buffer[ptr] != 0 {
					ptr = ptr + 2
				}
				ptr = ptr - 1
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 2
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 3
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 2
				for buffer[ptr] != 0 {
					ptr = ptr - 2
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 2
				for buffer[ptr] != 0 {
					ptr = ptr + 2
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 2
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 2
				for buffer[ptr] != 0 {
					ptr = ptr - 2
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 4
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 2
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 2
				for buffer[ptr] != 0 {
					ptr = ptr - 2
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 2
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr - 2
				for buffer[ptr] != 0 {
					ptr = ptr + 2
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 2
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 2
				for buffer[ptr] != 0 {
					ptr = ptr - 2
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 2
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr - 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 2
					for buffer[ptr] != 0 {
						ptr = ptr - 2
					}
					ptr = ptr - 2
					for buffer[ptr] != 0 {
						ptr = ptr - 2
					}
					ptr = ptr - 5
					buffer[ptr] = buffer[ptr] + byte(2)
					ptr = ptr + 2
				}
				ptr = ptr + 2
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 6
					for buffer[ptr] != 0 {
						ptr = ptr - 2
					}
					ptr = ptr - 2
					for buffer[ptr] != 0 {
						ptr = ptr - 2
					}
					ptr = ptr - 1
				}
				ptr = ptr + 1
			}
			ptr = ptr - 1
		}
		ptr = ptr + 1
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
			ptr = ptr + 2
			for buffer[ptr] != 0 {
				ptr = ptr + 2
			}
			ptr = ptr + 2
			for buffer[ptr] != 0 {
				ptr = ptr + 2
			}
			ptr = ptr - 1
			for buffer[ptr] != 0 {
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 2
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 3
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + 1
			for buffer[ptr] != 0 {
				for buffer[ptr] != 0 {
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 2
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 2
			}
			ptr = ptr + 2
			for buffer[ptr] != 0 {
				ptr = ptr - 2
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 4
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 2
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + 2
			for buffer[ptr] != 0 {
				ptr = ptr - 2
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 2
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr - 2
			for buffer[ptr] != 0 {
				ptr = ptr + 2
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 2
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr + 2
			for buffer[ptr] != 0 {
				ptr = ptr - 2
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 2
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr - 4
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr - 2
				for buffer[ptr] != 0 {
					ptr = ptr - 2
				}
				ptr = ptr - 2
				for buffer[ptr] != 0 {
					ptr = ptr - 2
				}
				ptr = ptr - 5
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 2
			}
			ptr = ptr + 2
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr - 6
				for buffer[ptr] != 0 {
					ptr = ptr - 2
				}
				ptr = ptr - 2
				for buffer[ptr] != 0 {
					ptr = ptr - 2
				}
				ptr = ptr - 1
			}
			ptr = ptr + 1
		}
		ptr = ptr + 1
		for buffer[ptr] != 0 {
			ptr = ptr - 1
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr + 1
			buffer[ptr] = buffer[ptr] - byte(1)
		}
		ptr = ptr - 6
		for buffer[ptr] != 0 {
			ptr = ptr + 2
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr - 2
			buffer[ptr] = buffer[ptr] - byte(1)
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 2
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 2
				buffer[ptr] = buffer[ptr] + byte(1)
				for buffer[ptr] != 0 {
					ptr = ptr + 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 2
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] - byte(1)
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						for buffer[ptr] != 0 {
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
							}
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 1
						}
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 3
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 3
						}
						ptr = ptr - 1
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 3
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 3
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 1
					}
					ptr = ptr - 4
					for buffer[ptr] != 0 {
						ptr = ptr + 2
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 2
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 2
				}
				ptr = ptr - 6
			}
			ptr = ptr + 2
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr - 2
				buffer[ptr] = buffer[ptr] + byte(1)
				for buffer[ptr] != 0 {
					ptr = ptr + 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 2
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] - byte(1)
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						for buffer[ptr] != 0 {
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
							}
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 1
						}
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 3
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 3
						}
						ptr = ptr - 1
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 3
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 3
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 1
					}
					ptr = ptr - 4
					for buffer[ptr] != 0 {
						ptr = ptr - 2
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 2
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr - 2
				}
			}
		}
		ptr = ptr + 7
	}
}
